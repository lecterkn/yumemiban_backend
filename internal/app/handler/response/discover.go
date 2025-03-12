package response

import (
	"time"

	"github.com/google/uuid"
)

type DiscoverLatestResponse struct {
	List []DiscoverResponse `json:"list" validate:"required"`
}

type DiscoverResponse struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	UserId    uuid.UUID `json:"userId" validate:"required"`
	Nickname  string    `json:"nickname" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Novel     string    `json:"novel" validate:"required"`
	Likes     int       `json:"likes" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}
