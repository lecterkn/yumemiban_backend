package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/output"
)

type PostRepository interface {
	Create(context.Context, *entity.PostEntity) error
	FindById(context.Context, uuid.UUID) (*entity.PostEntity, error)
	FindLatestByLastId(context.Context, *uuid.UUID) ([]output.DiscoverUsecaseQueryOutput, error)
	LikePost(context.Context, *entity.PostLikeEntity) error
}
