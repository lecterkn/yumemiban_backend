package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lecterkn/yumemiban_backend/internal/app/common"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler/response"
)

const (
	AUTHORIZATION_HEADER = "Authorization"
	AUTHORIZATION_PREFIX = "Bearer "
)

type JWTMiddleware struct {
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

// JWT認証を行うミドルウェア
func (m *JWTMiddleware) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// ヘッダーからアクセストークンを取得
		header := ctx.Request().Header.Get(AUTHORIZATION_HEADER)
		// トークンの形式を確認
		if !strings.HasPrefix(header, AUTHORIZATION_PREFIX) {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "不正な認証ヘッダー",
			})
		}
		// トークン部分のみ取得
		token := strings.TrimPrefix(header, AUTHORIZATION_PREFIX)
		// トークンを検証・復号化
		claims, err := common.DecodeToken(token)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "トークンの復号化に失敗しました",
			})
		}
		// トークンからID取得
		sub, err := claims.GetSubject()
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: "トークンからユーザーIDを取得できませんでした",
			})
		}
		// コンテキストにユーザーIDを設定
		ctx.Set("userId", sub)
		// エンドポイントに処理を引き渡す
		return next(ctx)
	}
}
