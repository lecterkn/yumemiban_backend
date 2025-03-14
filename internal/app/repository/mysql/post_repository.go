package mysql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/repository/mysql/model"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/output"
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
        INSERT INTO posts(id, user_id, nickname, title, content, novel, created_at, updated_at)
        VALUES(:id, :userId, :nickname, :title, :content, :novel, :createdAt, :updatedAt)
    `
	return RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExec(query, map[string]any{
			"id":        postEntity.Id[:],
			"userId":    postEntity.UserId[:],
			"nickname":  postEntity.Nickname,
			"title":     postEntity.Title,
			"content":   postEntity.Content,
			"novel":     postEntity.Novel,
			"createdAt": postEntity.CreatedAt,
			"updatedAt": postEntity.UpdatedAt,
		})
		return err
	})
}

// 最新の投稿をシーク法によるページネーションをして取得
func (r *PostRepositoryImpl) FindLatestByLastId(ctx context.Context, lastId *uuid.UUID) ([]output.DiscoverUsecaseQueryOutput, error) {
	query := `
        SELECT posts.id, posts.user_id, posts.nickname, posts.title, posts.content, posts.novel, posts.created_at, posts.updated_at, 
            COUNT(post_likes.user_id) as likes
        FROM posts
        LEFT JOIN post_likes
            ON post_likes.post_id = posts.id
    `
	if lastId != nil {
		query += "WHERE id > ?"
	}
	query += `
        GROUP BY posts.id
        ORDER BY id DESC
        LIMIT 50
    `
	queryOutputList := []output.DiscoverUsecaseQueryOutput{}
	err := RunInTx(ctx, r.database, func(tx *sqlx.Tx) error {
		models := []model.PostQueryModel{}
		if lastId != nil {
			err := tx.Select(&models, query, (*lastId)[:])
			if err != nil {
				return err
			}
		} else {
			err := tx.Select(&models, query)
			if err != nil {
				return err
			}
		}
		for _, queryModel := range models {
			queryOutput, err := r.toOutput(&queryModel)
			if err != nil {
				return err
			}
			queryOutputList = append(queryOutputList, *queryOutput)
		}
		return nil
	})
	return queryOutputList, err
}

// 投稿をIDから取得
func (r *PostRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entity.PostEntity, error) {
	query := `
        SELECT id, user_id, nickname, title, content, novel, created_at, updated_at
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

// 投稿にいいねを付与する
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
		Title:     postModel.Title,
		Content:   postModel.Content,
		Novel:     postModel.Novel,
		CreatedAt: postModel.CreatedAt,
		UpdatedAt: postModel.UpdatedAt,
	}, nil
}

func (r *PostRepositoryImpl) toOutput(queryModel *model.PostQueryModel) (*output.DiscoverUsecaseQueryOutput, error) {
	id, err := uuid.FromBytes(queryModel.Id)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.FromBytes(queryModel.UserId)
	if err != nil {
		return nil, err
	}
	return &output.DiscoverUsecaseQueryOutput{
		Id:        id,
		UserId:    userId,
		Nickname:  queryModel.Nickname,
		Title:     queryModel.Title,
		Content:   queryModel.Content,
		Novel:     queryModel.Novel,
		Likes:     queryModel.Likes,
		CreatedAt: queryModel.CreatedAt,
		UpdatedAt: queryModel.UpdatedAt,
	}, nil
}
