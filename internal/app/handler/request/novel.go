package request

type NovelGenerateRequest struct {
	Content string `json:"content" validate:"required"`
}
