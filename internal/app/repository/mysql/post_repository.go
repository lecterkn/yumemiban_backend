package mysql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/repository/mysql/model"
)

type PostRepositoryImpl struct {
	database *sqlx.DB
}

func NewPostRepositoryImpl(database *sqlx.DB) port.PostRepository {
	return &PostRepositoryImpl{
		database,
	}
}

func (r *PostRepositoryImpl) Create(ctx context.Context, postEntity *entity.PostEntity) error {
	query := `
        INSERT INTO posts(id, user_id, nickname, content, novel, created_at, updated_at)
        VALUES(:id, :userId, :nickname, :content, :novel, :createdAt, :updatedAt)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(query, map[string]any{
			"id":        postEntity.Id[:],
			"userId":    postEntity.UserId[:],
			"nickname":  postEntity.Nickname,
			"content":   postEntity.Content,
			"novel":     postEntity.Novel,
			"createdAt": postEntity.CreatedAt,
			"updatedAt": postEntity.UpdatedAt,
		})
		return err
	})
}

func (r *PostRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.PostEntity, error) {
	query := `
        SELECT id, user_id, nickname, content, novel, created_at, updated_at
        FROM posts
        WHERE id = ?
        LIMIT 1
    `
	postEntity := entity.PostEntity{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		postModel := model.PostModel{}
		err := tx.Get(&postModel, query, id[:])
		if err != nil {
			return err
		}
		postEntityPtr, err := r.toEntity(&postModel)
		if err != nil {
			return err
		}
		postEntity = *postEntityPtr
		return nil
	})
	return &postEntity, err
}

func (r *PostRepositoryImpl) LikePost(ctx context.Context, postLikeEntity *entity.PostLikeEntity) error {
	query := `
        INSERT INTO post_likes(post_id, user_id)
        VALUES(:postId, :userId)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(query, map[string]any{
			"postId": postLikeEntity.PostId[:],
			"userId": postLikeEntity.UserId[:],
		})
		return err
	})
}

func (r *PostRepositoryImpl) toEntity(postModel *model.PostModel) (*entity.PostEntity, error) {
	id, err := uuid.FromBytes(postModel.Id)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.FromBytes(postModel.UserId)
	if err != nil {
		return nil, err
	}
	return &entity.PostEntity{
		Id:        id,
		UserId:    userId,
		Nickname:  postModel.Nickname,
		Content:   postModel.Content,
		Novel:     postModel.Novel,
		CreatedAt: postModel.CreatedAt,
		UpdatedAt: postModel.UpdatedAt,
	}, nil
}
