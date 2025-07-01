import sys
from google import genai
from google.genai import types
from pathlib import Path



# GenAI クライアント初期化
client = genai.Client()


def run_deep_research(topic: str):
    """
    Google Search Grounding を使い Gemini 2.0 Flash でレポート生成し、
    grounding_metadata を合わせて返却。
    """
    grounding_tool = types.Tool(
        google_search=types.GoogleSearch()
    )
    config = types.GenerateContentConfig(
        temperature=0.7,
        max_output_tokens=1024,
        tools=[grounding_tool]
    )
    response = client.models.generate_content(
        model="gemini-2.5-flash",
        contents=[types.Content(
            role="user",
            parts=[types.Part(text=f"以下のトピックについて詳細なレポートを作成してください:\n{topic}")]
        )],
        config=config,
    )

    candidate = response.candidates[0]
    report = "".join(part.text for part in candidate.content.parts)

    gm = getattr(candidate, "grounding_metadata", None)
    queries = getattr(gm, "web_search_queries", []) if gm else []
    chunks  = getattr(gm, "grounding_chunks", []) if gm else []

    research = {
        "queries": queries,
        "results": [
            {
                "title": c.web.title if c.web else "",
                "uri":   c.web.uri   if c.web else "",
                "snippet": getattr(c, "snippet", "")
            }
            for c in chunks
        ]
    }
    return report, research


def merge_report_and_research(report: str, research: dict) -> str:
    """
    report: 生成されたレポート本文
    research: {
        "queries": List[str],
        "results": List[{"title": str, "uri": str, "snippet": str}]
    }
    の形式を想定

    戻り値:
      report→queries→results を組み合わせた単一の文字列
    """
    parts = []

    # 1) レポート本文
    parts.append(report.strip())

    # 2) 検索クエリ
    queries = research.get("queries", [])
    if queries:
        parts.append("\n=== Search Queries ===")
        for idx, q in enumerate(queries, start=1):
            parts.append(f"{idx}. {q}")

    # 3) 検索結果
    results = research.get("results", [])
    if results:
        parts.append("\n=== Search Results ===")
        for idx, res in enumerate(results, start=1):
            title   = res.get("title", "").strip()
            uri     = res.get("uri", "").strip()
            snippet = res.get("snippet", "").strip()

            parts.append(f"\n-- Result {idx} --")
            if title:
                parts.append(f"Title: {title}")
            if uri:
                parts.append(f"URI: {uri}")
            if snippet:
                parts.append(f"Snippet: {snippet}")

    # 結合して返却
    return "\n".join(parts)


def main():
    # トピック取得
    if len(sys.argv) >= 2:
        topic = sys.argv[1]
    else:
        topic = input("調査トピックを入力してください: ")

    # 実行スクリプトと同じフォルダをベースディレクトリに設定
    base_dir = Path(__file__).resolve().parent
    # ファイル名は第2引数、指定なければデフォルト
    filename = sys.argv[2] if len(sys.argv) >= 3 else "deep_research_result.txt"
    outfile = base_dir / filename

    try:
        report, research = run_deep_research(topic)

        # ファイルに書き込み
        with outfile.open("w", encoding="utf-8") as f:
            f.write(f"▶ 調査トピック: {topic}\n\n")
            f.write("✔ レポート生成完了:\n\n")
            f.write(report + "\n\n")
            f.write("=== Deep Research 結果 ===\n")
            if research["queries"]:
                f.write("■ 実行した検索クエリ:\n")
                for q in research["queries"]:
                    f.write(f"  - {q}\n")
            else:
                f.write("※ grounding_metadata が付与されませんでした。\n")
            if research["results"]:
                f.write("\n■ 取得したウェブ結果:\n")
                for idx, item in enumerate(research["results"], 1):
                    f.write(f"{idx}. {item['title']}\n")
                    f.write(f"   URL    : {item['uri']}\n")
                    if item["snippet"]:
                        f.write(f"   抜粋   : {item['snippet']}\n")
                    f.write("\n")

        print(f"結果をファイルに書き込みました: {outfile}")

    except Exception as e:
        print(f"[ERROR] Deep Research 実行中にエラー発生: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()
