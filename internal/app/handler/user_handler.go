package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lecterkn/yumemiban_backend/internal/app/common"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler/response"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase,
	}
}

// @summary	ユーザーを作成する
// @tags		user
// @produce	json
// @success	200		{object}	response.UserSignupResponse
// @router		/signin [post]
func (h *UserHandler) Create(ctx echo.Context) error {
	// ランダムなユーザー名とパスワードでユーザーを新規作成
	output, err := h.userUsecase.CreateUser(
		input.UserUsecaseCreateInput{
			Name:     common.GenerateUserName(),
			Password: common.GeneratePassword(),
		},
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.UserSignupResponse(*output))
}
