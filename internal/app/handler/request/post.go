package request

type PostCreateRequest struct {
	Nickname string `json:"nickname" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Novel    string `json:"novel" validate:"required"`
}
