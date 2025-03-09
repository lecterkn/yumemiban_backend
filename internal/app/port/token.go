package port

import (
	"github.com/google/uuid"
	"github.com/lecterkn/yumemiban_backend/internal/app/entity"
)

type TokenRepository interface {
	SaveRefreshToken(*entity.RefreshTokenEntity) error
	FindRefreshTokenByUserId(uuid.UUID) ([]entity.RefreshTokenEntity, error)
}
