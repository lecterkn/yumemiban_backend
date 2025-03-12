package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler/response"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/output"
)

type DiscoverHandler struct {
	discoverUsecase *usecase.DiscoverUsecase
}

func NewDiscoverHandler(discoverUsecase *usecase.DiscoverUsecase) *DiscoverHandler {
	return &DiscoverHandler{
		discoverUsecase,
	}
}

//	@summary		GetLatest
//	@description	最新の投稿を一覧取得する
//	@tags			discover
//	@produce		json
//	@param			lastId	query		string	false	"最後の投稿ID"
//	@success		200		{object}	response.DiscoverLatestResponse
//	@router			/discover/latest [get]
func (h *DiscoverHandler) Latest(ctx echo.Context) error {
	var lastId *uuid.UUID = nil
	// lastId取得
	if ctx.QueryParams().Has("lastId") {
		id, err := uuid.Parse(ctx.QueryParam("lastId"))
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Message: "lastIdの形式が不正です",
			})
		}
		lastId = &id
	}
	// 最新の投稿を一覧取得
	output, err := h.discoverUsecase.FindLatest(input.DiscoverUsecaseQueryInput{
		LastId: lastId,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, h.toResponse(output))
}

// Listのアウトプットをレスポンスに変換
func (h *DiscoverHandler) toResponse(queryListOutput *output.DiscoverUsecaseListQueryOutput) *response.DiscoverLatestResponse {
	list := []response.DiscoverResponse{}
	for _, output := range queryListOutput.List {
		list = append(list, response.DiscoverResponse(output))
	}
	return &response.DiscoverLatestResponse{
		List: list,
	}
}
