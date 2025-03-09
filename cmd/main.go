package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/lecterkn/yumemiban_backend/internal/app/di"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".envファイルが存在しません")
	}
	port, ok := os.LookupEnv("ECHO_SERVER_PORT")
	if !ok {
		panic("環境変数\"ECHO_SERVER_PORT\"が設定されていません")
	}
	app := echo.New()
	setRouting(app)
	app.Logger.Fatal(app.Start(":" + port))
}

// エンドポイントのルーティング
func setRouting(app *echo.Echo) {
	handlerSet := di.InitializeHandlerSet()

	// /api
	api := app.Group("api")

	// /api/signup
	api.POST("/signup", handlerSet.UserHandler.Create)

	// 認証対象グループ
	auth := app.Group("")
	auth.Use(handlerSet.JWTMiddleware.Authorization)
}
