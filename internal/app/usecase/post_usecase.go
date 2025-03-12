package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/output"
)

type PostUsecase struct {
	txProvider     port.TransactionProvider
	postRepository port.PostRepository
	userRepository port.UserRepository
}

func NewPostUsecase(
	txProvider port.TransactionProvider,
	postRepository port.PostRepository,
	userRepository port.UserRepository,
) *PostUsecase {
	return &PostUsecase{
		txProvider,
		postRepository,
		userRepository,
	}
}

func (u *PostUsecase) CreatePost(userId uuid.UUID, cmd input.PostUsecaseCreateInput) (*output.PostUsecaseCreateOutput, error) {
	postOutput := output.PostUsecaseCreateOutput{}
	// トランザクション
	err := u.txProvider.Transact(func(ctx context.Context) error {
		// ユーザー存在確認
		userEntity, err := u.userRepository.FindById(ctx, userId)
		if err != nil {
			return err
		}
		// 投稿作成
		postEntity, err := entity.NewPostEntity(userEntity.Id, cmd.Nickname, cmd.Content, cmd.Novel)
		if err != nil {
			return err
		}
		// リポジトリで投稿作成
		err = u.postRepository.Create(ctx, postEntity)
		if err != nil {
			return err
		}
		postOutput = output.PostUsecaseCreateOutput(*postEntity)
		return nil
	})
	return &postOutput, err
}

func (u *PostUsecase) LikePost(postId, userId uuid.UUID) error {
	return u.txProvider.Transact(func(ctx context.Context) error {
		// 投稿存在確認
		_, err := u.postRepository.FindById(ctx, postId)
		if err != nil {
			return err
		}
		// ユーザー存在確認
		_, err = u.userRepository.FindById(ctx, userId)
		if err != nil {
			return err
		}
		// いいね作成
		postLikeEntity := entity.NewPostLikeEntity(postId, userId)
		return u.postRepository.LikePost(ctx, postLikeEntity)
	})
}
