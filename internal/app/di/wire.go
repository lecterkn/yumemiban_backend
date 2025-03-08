//go:build wireinject
// +build wireinject

package di

import "github.com/google/wire"

var databaseSet = wire.NewSet(
// データベースのコネクタ
)

var repositorySet = wire.NewSet(
// リポジトリの実装
)

var usecaseSet = wire.NewSet(
// ユースケース
)

var handlerSet = wire.NewSet(
// ハンドラ
)

// 生成されるハンドラ
type HandlerSet struct {
}

// ハンドラセットを取得する
func InitializeControllerSet() *HandlerSet {
	wire.Build(
		databaseSet,
		repositorySet,
		usecaseSet,
		handlerSet,
		wire.Struct(new(HandlerSet), "*"),
	)
	return nil
}
