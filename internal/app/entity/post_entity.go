package entity

import (
	"time"

	"github.com/google/uuid"
)

type PostEntity struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Nickname  string
	Title     string
	Content   string
	Novel     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPostEntity(userId uuid.UUID, nickname, title, content, novel string) (*PostEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &PostEntity{
		Id:        id,
		UserId:    userId,
		Nickname:  nickname,
		Title:     title,
		Content:   content,
		Novel:     novel,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

type PostLikeEntity struct {
	PostId uuid.UUID
	UserId uuid.UUID
}

func NewPostLikeEntity(postId, userId uuid.UUID) *PostLikeEntity {
	return &PostLikeEntity{
		PostId: postId,
		UserId: userId,
	}
}
