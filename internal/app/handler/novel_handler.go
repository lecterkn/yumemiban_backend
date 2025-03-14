package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler/request"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler/response"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
)

type NovelHandler struct {
	novelUsecase *usecase.NovelUsecase
}

func NewNovelHandler(novelUsecase *usecase.NovelUsecase) *NovelHandler {
	return &NovelHandler{
		novelUsecase,
	}
}

//	@summary		GenerateNovel
//	@description	小説を生成する
//	@tags			novel
//	@produce		json
//	@param			request	body		request.NovelGenerateRequest	true	"小説生成リクエスト"
//	@success		200		{object}	response.NovelGenerateResponse
//	@router			/novels/generate [post]
//	@security		BearerAuth
func (h *NovelHandler) Generate(ctx echo.Context) error {
	novelRequest := request.NovelGenerateRequest{}
	// リクエスト取得
	if err := ctx.Bind(&novelRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "リクエストの形式が不正です",
		})
	}
	// 小説を生成
	output, err := h.novelUsecase.GenerateNovel(input.NovelGenerateInput(novelRequest))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.NovelGenerateResponse(*output))
}
