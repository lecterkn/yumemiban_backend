package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
)

type PostRepository interface {
	Create(context.Context, *entity.PostEntity) error
	FindById(context.Context, uuid.UUID) (*entity.PostEntity, error)
	LikePost(context.Context, *entity.PostLikeEntity) error
}
