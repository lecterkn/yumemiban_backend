package model

import "time"

type PostModel struct {
	Id        []byte    `db:"id"`
	UserId    []byte    `db:"user_id"`
	Nickname  string    `db:"nickname"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Novel     string    `db:"novel"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PostLikeModel struct {
	PostId []byte `db:"post_id"`
	UserId []byte `db:"user_id"`
}

type PostQueryModel struct {
	Id        []byte    `db:"id"`
	UserId    []byte    `db:"user_id"`
	Nickname  string    `db:"nickname"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Novel     string    `db:"novel"`
	Likes     int       `db:"likes"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
