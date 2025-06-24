#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
auto_deep_research.py

Google GenAI クライアントを使って Gemini 2.0 Flash の
“google_search” Grounding 機能を呼び出し、自動検索＋レポート生成を行うサンプル。

## 準備
1. `pip install google-genai`
2. 環境変数に GOOGLE_APPLICATION_CREDENTIALS をセットしておく
3. GCP プロジェクトIDとリージョンは以下に直接設定

## 使い方
```bash
# 引数あり
$ python auto_deep_research.py "調査トピック"

# 引数なし（対話形式）
$ python auto_deep_research.py
```"""
import sys
from google import genai
from google.genai import types

# ─── 設定 ───
PROJECT_ID = "auto-deepresearch"  # GCP プロジェクトID
LOCATION   = "us-central1"        # GCP リージョン

# GenAI クライアント初期化
# Vertex AI (Google Cloud API) を利用する場合は vertexai=True, project, location を指定
client = genai.Client(vertexai=True, project=PROJECT_ID, location=LOCATION)


def run_deep_research(topic: str) -> str:
    """
    Google Search Grounding を使い Gemini 2.0 Flash でレポート生成
    """
    # Grounding Tool の定義
    grounding_tool = types.Tool(
        google_search=types.GoogleSearch()
    )
    # 生成設定
    config = types.GenerateContentConfig(
        temperature=0.7,
        max_output_tokens=1024,
        tools=[grounding_tool]
    )
    # コンテンツ呼び出し
    response = client.models.generate_content(
        model="gemini-2.0-flash",
        contents=[
            types.Content(
                role="user",
                parts=[
                    types.Part(text=f"以下のトピックについて詳細なレポートを作成してください:\n{topic}")
                ]
            )
        ],
        config=config,
    )
    return response.text


def main():
    if len(sys.argv) >= 2:
        topic = sys.argv[1]
    else:
        topic = input("調査トピックを入力してください: ")

    print(f"▶ 調査トピック: {topic}\n")
    try:
        result = run_deep_research(topic)
        print("✔ レポート生成完了:\n")
        print(result)
    except Exception as e:
        print(f"[ERROR] Deep Research 実行中にエラー発生: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()

