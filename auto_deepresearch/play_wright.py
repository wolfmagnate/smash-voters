import time
from playwright.sync_api import sync_playwright, TimeoutError

# --- 設定項目 ---
# ステップ2で調べたパスに書き換えてください
# Windowsの場合: r"C:\Users\YourUser\AppData\Local\Google\Chrome\User Data"
# Macの場合: "/Users/YourUser/Library/Application Support/Google/Chrome"
CHROME_USER_DATA_PATH = r"C:/Users/sohei/AppData/Local/Google/Chrome/User Data"

PROMPT = "日本の伝統的なお祭りについて、その歴史的背景と現代における意義を教えてください。"
# --- 設定項目ここまで ---

with sync_playwright() as p:
    browser = None
    try:
        # 既存のユーザープロファイルを使用してブラウザを起動
        context = p.chromium.launch_persistent_context(
            user_data_dir=CHROME_USER_DATA_PATH,
            headless=False,  # Falseでブラウザを表示、Trueで非表示
            channel="chrome", # 通常インストールされたChromeを使用
            args=['--profile-directory=Default_profile']
        )
        page = context.new_page()
        page.set_default_timeout(20000) # タイムアウトを20秒に設定

        # Geminiのページにアクセス
        page.goto("https://gemini.google.com/", wait_until="networkidle")
        print("Geminiのページにアクセスしました。")

        

    except TimeoutError:
        print("エラー: タイムアウトしました。要素が見つからないか、ページの読み込みが遅い可能性があります。")
    except Exception as e:
        print(f"予期せぬエラーが発生しました: {e}")
    finally:
        if browser:
            # ブラウザを閉じる場合は以下のコメントを外す
            # input("何かキーを押すとブラウザを閉じます...")
            # context.close()
            pass