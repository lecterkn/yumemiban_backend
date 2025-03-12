package usecase

import (
	"context"

	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/output"
)

type DiscoverUsecase struct {
	postRepository port.PostRepository
}

func NewDiscoverUsecase(postRepository port.PostRepository) *DiscoverUsecase {
	return &DiscoverUsecase{
		postRepository,
	}
}

// 最新のものから投稿を一覧取得
func (u *DiscoverUsecase) FindLatest(cmd input.DiscoverUsecaseQueryInput) (*output.DiscoverUsecaseListQueryOutput, error) {
	outputList, err := u.postRepository.FindLatestByLastId(context.Background(), cmd.LastId)
	if err != nil {
		return nil, err
	}
	return &output.DiscoverUsecaseListQueryOutput{
		List: outputList,
	}, nil
}
