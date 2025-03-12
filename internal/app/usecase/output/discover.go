package output

import (
	"time"

	"github.com/google/uuid"
)

type DiscoverUsecaseListQueryOutput struct {
	List []DiscoverUsecaseQueryOutput
}

type DiscoverUsecaseQueryOutput struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Nickname  string
	Content   string
	Novel     string
	Likes     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
