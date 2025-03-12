//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/lecterkn/yumemiban_backend/internal/app/database"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler"
	"github.com/lecterkn/yumemiban_backend/internal/app/provider"
	"github.com/lecterkn/yumemiban_backend/internal/app/repository/mysql"
	"github.com/lecterkn/yumemiban_backend/internal/app/repository/redis"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase"
)

// データベースのコネクタ
var databaseSet = wire.NewSet(
	database.GetMySQLConnection,
	database.GetRedisClient,
)

// リポジトリの実装
var repositorySet = wire.NewSet(
	mysql.NewUserRepositoryImpl,
	redis.NewTokenRepositoryImpl,
)

// プロバイダの実装
var providerSet = wire.NewSet(
	provider.NewTransactionProviderImpl,
)

// ユースケース
var usecaseSet = wire.NewSet(
	usecase.NewUserUsecase,
)

// ハンドラ
var handlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewJWTMiddleware,
)

// 生成されるハンドラ
type HandlerSet struct {
	UserHandler   *handler.UserHandler
	JWTMiddleware *handler.JWTMiddleware
}

// ハンドラセットを取得する
func InitializeHandlerSet() *HandlerSet {
	wire.Build(
		databaseSet,
		repositorySet,
		providerSet,
		usecaseSet,
		handlerSet,
		wire.Struct(new(HandlerSet), "*"),
	)
	return nil
}
