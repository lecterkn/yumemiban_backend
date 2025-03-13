package output

import (
	"time"

	"github.com/google/uuid"
)

type PostUsecaseCreateOutput struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Nickname  string
	Title     string
	Content   string
	Novel     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
