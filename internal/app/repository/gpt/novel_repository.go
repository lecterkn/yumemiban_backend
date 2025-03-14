package gpt

import (
	"context"

	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/sashabaranov/go-openai"
)

const (
	MAX_TOKEN = 127
)

type NovelRepositoryImpl struct {
	gptClient *openai.Client
}

func NewNovelRepositoryImpl(gptClient *openai.Client) port.NovelRepository {
	return &NovelRepositoryImpl{
		gptClient,
	}
}

// 小説を生成する
func (r *NovelRepositoryImpl) GenerateNovel(content string) (*string, error) {
	response, err := r.gptClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT4oMini,
			MaxTokens: MAX_TOKEN,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: r.getSystemPrompt(),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return &response.Choices[0].Message.Content, nil
}

// システムプロンプト
func (r *NovelRepositoryImpl) getSystemPrompt() string {
	return `
        あなたは短編小説家です
        夢の内容を一つの短い物語に作り替えます
        あなたは100文字程度の簡潔な小説を執筆します
    `
}
