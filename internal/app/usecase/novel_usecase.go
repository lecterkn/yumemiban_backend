package usecase

import (
	"errors"

	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/input"
	"github.com/lecterkn/yumemiban_backend/internal/app/usecase/output"
)

type NovelUsecase struct {
	novelRepository port.NovelRepository
}

func NewNovelUsecase(novelRepository port.NovelRepository) *NovelUsecase {
	return &NovelUsecase{
		novelRepository,
	}
}

func (u *NovelUsecase) GenerateNovel(cmd input.NovelGenerateInput) (*output.NovelGenerateOutput, error) {
	if len(cmd.Content) > 100 {
		return nil, errors.New("文字数が100文字以上です")
	}
	novel, err := u.novelRepository.GenerateNovel(cmd.Content)
	if err != nil {
		return nil, err
	}
	return &output.NovelGenerateOutput{
		Novel: *novel,
	}, nil
}
