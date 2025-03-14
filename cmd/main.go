package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lecterkn/yumemiban_backend/docs"
	"github.com/lecterkn/yumemiban_backend/internal/app/di"
	echoSwagger "github.com/swaggo/echo-swagger"
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
//
//	@title						yumemi backend API
//	@version					1.0
//	@description				YumemibanのAPIサーバー
//	@host						http://localhost:8089
//	@BasePath					/api
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
func setRouting(app *echo.Echo) {
	// CORS
	app.Use(middleware.CORS())

	// swagger
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	handlerSet := di.InitializeHandlerSet()

	// /api
	api := app.Group("api")

	// /api/signup
	api.POST("/signup", handlerSet.UserHandler.Create)

	// /api/discover/latest
	api.GET("/discover/latest", handlerSet.DiscoverHandler.Latest)

	// 認証対象グループ
	auth := api.Group("")
	auth.Use(handlerSet.JWTMiddleware.Authorization)

	auth.POST("/posts", handlerSet.PostHandler.Create)
	api.GET("/posts/:postId", handlerSet.PostHandler.Get)
	auth.POST("/posts/:postId/likes", handlerSet.PostHandler.Like)

	auth.POST("/novels/generate", handlerSet.NovelHandler.Generate)
}
