package input

type NovelGenerateInput struct {
	Content string `json:"content" validate:"required"`
}
