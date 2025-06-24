package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wolfmagnate/smash-voters/coverage/handlers"
	"github.com/wolfmagnate/smash-voters/coverage/services"
)

func main() {
	// Initialize other services
	researchService := services.NewResearchService("http://localhost:8000/research")

	// Initialize handlers
	researchHandler := handlers.NewResearchHandler(researchService)

	// Echoのインスタンスを作成
	e := echo.New()

	// ミドルウェアを設定（推奨）
	// リクエストID、ロガー、パニックからの復帰など
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// ヘルスチェック
	e.GET("/healthz", func(c echo.Context) error {
		// c.String(ステータスコード, 返す文字列)
		return c.String(http.StatusOK, "healthy")
	})

	// API endpoints
	api := e.Group("/api/v1")

	// Research API endpoint
	api.POST("/research/:theme/:is_positive", researchHandler.HandleResearch)

	// Webサーバーをポート8080で起動
	// e.Logger.Fatal はエラーが発生した場合にログを出力してプログラムを終了します
	e.Logger.Fatal(e.Start(":8080"))
}
