package response

import (
	"time"

	"github.com/google/uuid"
)

type PostCreateResponse struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	UserId    uuid.UUID `json:"userId" validate:"required"`
	Nickname  string    `json:"nickname" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Novel     string    `json:"novel" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}
