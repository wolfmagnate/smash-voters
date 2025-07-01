# server.py
from fastapi import FastAPI, BackgroundTasks, HTTPException, Request
from fastapi.responses import JSONResponse
from pydantic import BaseModel, HttpUrl, Field
import uuid
import requests
import os
import io
from aut_websearch_update import run_deep_research, merge_report_and_research
from googleapiclient.http import MediaFileUpload
from googleapiclient.errors import HttpError
from google.cloud import storage

credentials_path = 'service_account_key.json' 
os.environ['GOOGLE_APPLICATION_CREDENTIALS'] = credentials_path

# アクセスしたいGCSバケット名
BUCKET_NAME = 'smash-voters-report' # ここをあなたのバケット名に変更してください

app = FastAPI(
    title="Deep Research API",
    description="時間のかかる調査タスクを非同期で実行するAPI。ジョブIDを返し、完了後にWebhookへ通知します。",
    version="1.0.0",
    openapi_url="/v1/openapi.json",
    docs_url="/v1/docs"
)

# --- Pydanticモデル定義 ---
class ResearchRequest(BaseModel):
    query: str = Field(..., description="調査したい内容のクエリ")
    drive_path: str = Field(..., description="結果レポートを保存するGoogle Driveなどのパス")
    webhook_url: HttpUrl = Field(..., description="処理完了時に通知を送るWebhook URL")

class Job(BaseModel):
    job_id: str = Field(..., description="非同期処理を追跡するための一意なジョブID")

class ApiError(BaseModel):
    code: str
    message: str

class WebhookPayload(BaseModel):
    job_id: str
    status: str
    result_path: str

# バリデーションエラーを ApiError で返す
#リクエストの型が一致しないときにこれを返す
from fastapi.exceptions import RequestValidationError
#もし RequestValidationError というエラーが発生したら、次に書く関数を使って処理してね
@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request: Request, exc: RequestValidationError):
    return JSONResponse(
        status_code=400, 
        content={
        "code": "INVALID_PARAMETER",
        "message": str(exc)
    })
#エンドポイントをv1/researchとしてpostメソッドを受け取る
@app.post(
    "/v1/research",
    response_model=Job,
    status_code=202,
    responses={
        400: {"model": ApiError},#このエラーは上の400エラーとは異なるのか
        202: {"model": Job}
    }
)

#process_research’ という別の関数を、裏側（バックグラウンド）で動かすように登録
#世界でひとつだけのランダムな文字列（UUID）を作って、ジョブIDとして文字列に変えます
async def start_research_job(request: ResearchRequest, background_tasks: BackgroundTasks):
    job_id = str(uuid.uuid4())
    background_tasks.add_task(process_research, job_id, request)
    return Job(job_id=job_id)

async def process_research(job_id: str, req: ResearchRequest):
    filename = f"{job_id}.txt"
    #run_deep_research’ 関数を呼んで、調査結果レポートと生のリサーチデータを取得し,取得したレポートとリサーチデータをくっつけて、ひとつの文章にまとめて変数格納
    try:
        # ① Deep Research 実行．
        print("start research")
        report, research = run_deep_research(req.query)
        result = merge_report_and_research(report, research)
        
        print("start save")
        save_text_to_gcs(BUCKET_NAME,req.drive_path,result)
        status = 'completed'
    except Exception:
        result_path = ""
        status = 'failed'

    # Webhook へ通知
    payload = WebhookPayload(job_id=job_id, status=status, result_path=req.drive_path)
    try:
        print("call payload")
        requests.post(req.webhook_url, json=payload.model_dump())
        print("call finished")
    except Exception:
        pass


    

def save_text_to_gcs(bucket_name, destination_blob_name, text_content):
    """
    文字列をGCSバケット内の指定されたパスにテキストファイルとして保存する関数
    """
    try:
        # GCSクライアントを初期化
        storage_client = storage.Client()
        print("get client")

        # バケットを取得
        bucket = storage_client.bucket(bucket_name)
        print("get bucket")

        # 保存先のオブジェクト（blob）を指定
        blob = bucket.blob(destination_blob_name)
        
        print(f"🔄 ファイルをアップロードしています...")
        print(f"  - バケット: {bucket_name}")
        print(f"  - 保存先: {destination_blob_name}")

        # 文字列をUTF-8としてアップロード
        blob.upload_from_string(text_content, content_type='text/plain; charset=utf-8')

        print(f"✅ ファイルの保存が完了しました。")
        print(f"   gs://{bucket_name}/{destination_blob_name}")

    except Exception as e:
        print(f"🚨 エラーが発生しました: {e}")
        print("ヒント: バケット名や認証キーのパスが正しいか、権限が付与されているか確認してください。")