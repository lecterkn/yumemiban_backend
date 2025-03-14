package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler/request"
	"github.com/lecterkn/yumemiban_backend/internal/app/handler/response"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
)

type PostHandler struct {
	postUsecase *usecase.PostUsecase
}

func NewPostHandler(postUsecase *usecase.PostUsecase) *PostHandler {
	return &PostHandler{
		postUsecase,
	}
}

// @summary		CreatePost
// @description	投稿を新規作成する
// @tags			post
// @produce		json
// @param			request	body		request.PostCreateRequest	true	"投稿作成リクエスト"
// @success		200		{object}	response.PostCreateResponse
// @router			/posts [post]
func (h *PostHandler) Create(ctx echo.Context) error {
	// 投稿作成リクエスト取得
	postCreateRequest := request.PostCreateRequest{}
	if err := ctx.Bind(&postCreateRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "リクエストのボディが不正です",
		})
	}
	// ユーザーID取得
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	// 投稿作成
	output, err := h.postUsecase.CreatePost(
		userId,
		input.PostUsecaseCreateInput{
			Nickname: postCreateRequest.Nickname,
			Title:    postCreateRequest.Title,
			Content:  postCreateRequest.Content,
			Novel:    postCreateRequest.Novel,
		},
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.PostCreateResponse(*output))
}

// @summary		LikePost
// @description	投稿にいいねをつける
// @tags			post
// @param			postId	path	string	true	"投稿ID"
// @success		204
// @router			/posts/{postId}/likes [post]
func (h *PostHandler) Like(ctx echo.Context) error {
	// ユーザーID取得
	userId, err := uuid.Parse(ctx.Get("userId").(string))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	postId, err := uuid.Parse(ctx.Param("postId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "投稿IDが不正です",
		})
	}
	err = h.postUsecase.LikePost(postId, userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}
