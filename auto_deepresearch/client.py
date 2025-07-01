# client_test.py
import threading
import time
from fastapi import FastAPI, Request
import uvicorn
import requests

# ─── テスト用Webhookサーバー ───
#これは「FastAPI」というものを準備して、webhook_app という箱に入れているイメージ
#どのurlにどのhttpメソッドでアクセスが来たときにどの関数を呼び出すかを登録していく
webhook_app = FastAPI()

#以下のurlにpostリクエストが来たらそれ以下の関数を実行する
#request（リクエスト）には、送られてきたデータが入っています(どんな形を受け付けるのかはtest_req?)
#request.json() は、送られてきたデータを「JSON」というかたちから、Pythonの辞書（dict）に変える
@webhook_app.post("/webhook/research-complete")
async def receive_webhook(request: Request):
    payload = await request.json()
    print("=== Webhook 受信 ===")
    print(payload)
    return {}  # 200 OK

def run_test():
    # サーバー起動待機
    time.sleep(1)
    test_req = {
        "query": "AI市場調査",
        "drive_path": "reports/file2.txt",
        "webhook_url": "http://localhost:8001/webhook/research-complete"
    }
    #送る先は http://localhost:8000/v1/research というアドレス（動いている自分のサーバー．どこにあるか確認）
    resp = requests.post("http://localhost:8000/v1/research", json=test_req)
    print("=== POST /v1/research レスポンス ===")
    print(resp.status_code, resp.json())

if __name__ == '__main__':
    #webhookの8001番ポートサーバを別スレッドで立てて接続を維持したまま，メインスレッドでrun_testを実施する
    #8001番を立てているが，ここにリクエストは実際は送っていないよな．．
    threading.Thread(
        target=lambda: uvicorn.run(webhook_app, host='localhost', port=8001),
        daemon=True
    ).start()
    # テスト実行
    run_test()
    # Webhook の受信を少し待つ
    time.sleep(1000)
