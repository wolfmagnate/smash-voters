#どのファイルがどの役割をしていて，どのようにコネクトするかを明確にして作る
import os
from pathlib import Path
from typing import Optional

from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

# 他モジュールから機能をインポート
from aut_websearch_update import run_deep_research  # レポート生成ロジック
from google_drive_test import upload_to_gdoc      # Driveアップロードロジック

# FastAPI アプリケーション定義
app = FastAPI()

# リクエストボディ用スキーマ
class ResearchRequest(BaseModel):
    folder_id: str     # Google Drive 上の親フォルダID
    topic: str         # 調査トピック（クエリ文字列）
    filename: str      # 出力ファイル名（例: report.txt）

# レスポンス用スキーマ
class ResearchResponse(BaseModel):
    status: str
    file_id: Optional[str] = None
    web_view_link: Optional[str] = None
    detail: Optional[str] = None

@app.post("/aut_websearch_update", response_model=ResearchResponse)
async def deep_research_endpoint(req: ResearchRequest):
    # ファイル名にパス区切り文字が含まれていないか確認
    if any(sep in req.filename for sep in ("/", "\\")):
        raise HTTPException(status_code=400, detail="不正なファイル名です。パス区切り文字は禁止されています。")

    # 実行スクリプトと同じフォルダに出力
    base_dir = Path(__file__).resolve().parent
    outfile = base_dir / req.filename

    try:
        # 1) レポート生成（deep_research モジュール）
        report, research = run_deep_research(req.topic)

        # 2) テキストファイルに書き込み
        with outfile.open("w", encoding="utf-8") as f:
            f.write(f"▶ 調査トピック: {req.topic}\n\n")
            f.write("✔ レポート生成完了:\n\n")
            f.write(report + "\n\n")
            f.write("=== Deep Research 結果 ===\n")
            if research.get("queries"):
                f.write("■ 実行した検索クエリ:\n")
                for q in research["queries"]:
                    f.write(f"  - {q}\n")
            else:
                f.write("※ grounding_metadata が付与されませんでした。\n")
            if research.get("results"):
                f.write("\n■ 取得したウェブ結果:\n")
                for idx, item in enumerate(research["results"], start=1):
                    f.write(f"{idx}. {item['title']}\n")
                    f.write(f"   URL    : {item['uri']}\n")
                    if item.get('snippet'):
                        f.write(f"   抜粋   : {item['snippet']}\n")
                    f.write("\n")

        # 3) Driveアップロード（drive_upload モジュール）
        file_id, web_link = upload_to_gdoc(outfile, req.folder_id, req.filename)
        return ResearchResponse(status="success", file_id=file_id, web_view_link=web_link)

    except HTTPException:
        # FastAPI が自動でハンドル
        raise
    except Exception as e:
        # その他想定外エラー
        raise HTTPException(status_code=500, detail=str(e))

# エントリポイント
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
