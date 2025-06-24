package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echoのインスタンスを作成
	e := echo.New()

	// ミドルウェアを設定（推奨）
	// リクエストID、ロガー、パニックからの復帰など
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ヘルスチェック
	e.GET("/healthz", func(c echo.Context) error {
		// c.String(ステータスコード, 返す文字列)
		return c.String(http.StatusOK, "healthy")
	})

	// Webサーバーをポート8080で起動
	// e.Logger.Fatal はエラーが発生した場合にログを出力してプログラムを終了します
	e.Logger.Fatal(e.Start(":8080"))
}
