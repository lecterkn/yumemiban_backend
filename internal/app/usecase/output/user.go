package output

import (
	"time"

	"github.com/google/uuid"
)

type UserUsecaseCreateOutput struct {
	Id           uuid.UUID
	Name         string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	AccessToken  string
	RefreshToken string
}
