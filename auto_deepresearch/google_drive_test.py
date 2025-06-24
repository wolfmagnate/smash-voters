import os
from google.oauth2 import service_account
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError

# --- 設定項目 ---
# 1. サービスアカウントの秘密鍵ファイルへのパス
SERVICE_ACCOUNT_FILE = 'service_account_key.json' 
# 2. Google Drive APIのスコープ
SCOPES = ['https://www.googleapis.com/auth/drive']
# 3. 操作したい親フォルダのID
PARENT_FOLDER_ID = '1yd7VBM4LjTajzGWBnkhgSfXK7TjLQt1b' # ここに共有設定したフォルダのIDを指定

def main():
    """
    サービスアカウント認証を使用し、手動の同意なしでGoogle Driveにアクセスする。
    """
    try:
        # サービスアカウントの認証情報を作成
        creds = service_account.Credentials.from_service_account_file(
            SERVICE_ACCOUNT_FILE, scopes=SCOPES)

        # Google Drive APIのサービスオブジェクトを構築
        service = build('drive', 'v3', credentials=creds)

        # --- これ以降のDrive操作はこれまでと同じ ---

        # 例：フォルダ内に'Automated Document'という名前のファイルを作成
        file_metadata = {
            'name': 'Automated Document',
            'mimeType': 'application/vnd.google-apps.document',
            'parents': [PARENT_FOLDER_ID]
        }
        
        # 共有ドライブの場合は supportsAllDrives=True を追加
        file = service.files().create(
            body=file_metadata, 
            fields='id, name, webViewLink'
            # supportsAllDrives=True 
        ).execute()

        print(f"ファイルを作成しました。")
        print(f"ファイル名: {file.get('name')}")
        print(f"ファイルID: {file.get('id')}")
        print(f"リンク: {file.get('webViewLink')}")

    except FileNotFoundError:
        print(f"エラー: 秘密鍵ファイル '{SERVICE_ACCOUNT_FILE}' が見つかりません。")
    except HttpError as error:
        print(f"APIエラーが発生しました: {error}")
    except Exception as e:
        print(f"予期せぬエラーが発生しました: {e}")


if __name__ == '__main__':
    main()