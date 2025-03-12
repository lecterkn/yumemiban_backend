package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lecterkn/yumemiban_backend/internal/app/common"
)

type UserEntity struct {
	Id        uuid.UUID
	Name      string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 新規ユーザーを作成
func NewUserEntity(name, password string) (*UserEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	pass, err := common.EncryptPassword(password)
	if err != nil {
		return nil, err
	}
	// TODO バリデーション
	return &UserEntity{
		Id:        id,
		Name:      name,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
