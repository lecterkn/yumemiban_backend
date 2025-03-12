package port

import (
	"context"

	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
)

type PostRepository interface {
	Create(context.Context, *entity.PostEntity) error
}
