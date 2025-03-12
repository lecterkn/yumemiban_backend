package usecase

import (
	"context"

	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/output"
)

type UserUsecase struct {
	txProvider      port.TransactionProvider
	userRepository  port.UserRepository
	tokenRepository port.TokenRepository
}

func NewUserUsecase(
	txProvider port.TransactionProvider,
	userRepository port.UserRepository,
	tokenRepository port.TokenRepository,
) *UserUsecase {
	return &UserUsecase{
		txProvider,
		userRepository,
		tokenRepository,
	}
}

// ユーザーを作成しトークンを生成する
func (u *UserUsecase) CreateUser(cmd input.UserUsecaseCreateInput) (*output.UserUsecaseCreateOutput, error) {
	userOutput := output.UserUsecaseCreateOutput{}
	// トランザクション開始
	err := u.txProvider.Transact(func(ctx context.Context) error {
		// ユーザー作成
		userEntity, err := entity.NewUserEntity(cmd.Name, cmd.Password)
		if err != nil {
			return err
		}
		// リポジトリでユーザーを作成
		err = u.userRepository.Create(ctx, userEntity)
		if err != nil {
			return err
		}
		// リフレッシュトークン生成
		refreshTokenEntity, err := entity.NewRefreshTokenEntity(userEntity.Id)
		if err != nil {
			return err
		}
		// アクセストークン生成
		accessTokenEntity, err := entity.NewAccessTokenEntity(userEntity.Id)
		if err != nil {
			return err
		}
		// リポジトリでリフレッシュトークンを保存
		err = u.tokenRepository.SaveRefreshToken(refreshTokenEntity)
		if err != nil {
			return err
		}
		userOutput = output.UserUsecaseCreateOutput{
			Id:           userEntity.Id,
			Name:         userEntity.Name,
			Password:     cmd.Password,
			CreatedAt:    userEntity.CreatedAt,
			UpdatedAt:    userEntity.UpdatedAt,
			AccessToken:  accessTokenEntity.Token,
			RefreshToken: refreshTokenEntity.Token,
		}
		return nil
	})
	return &userOutput, err
}
