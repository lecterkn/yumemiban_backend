package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
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
