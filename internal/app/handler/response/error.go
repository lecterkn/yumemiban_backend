package response

type ErrorResponse struct {
	Message string `json:"message" validate:"required"`
}
