package response

type NovelGenerateResponse struct {
	Novel string `json:"novel" validate:"required"`
}
