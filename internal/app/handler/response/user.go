package response

import (
	"time"

	"github.com/google/uuid"
)

type UserSignupResponse struct {
	Id           uuid.UUID `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Password     string    `json:"password" validate:"required"`
	CreatedAt    time.Time `json:"createdAt" validate:"required"`
	UpdatedAt    time.Time `json:"updatedAt" validate:"required"`
	AccessToken  string    `json:"accessToken" validate:"required"`
	RefreshToken string    `json:"refreshToken" validate:"required"`
}
