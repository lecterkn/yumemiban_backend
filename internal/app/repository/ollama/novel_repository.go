package ollama

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/lecterkn/yumemiban_backend/internal/app/database"
	"github.com/lecterkn/yumemiban_backend/internal/app/port"
	"github.com/lecterkn/yumemiban_backend/internal/app/repository/ollama/model"
)

type NovelRepositoryImpl struct {
	ollamaClient *database.OllamaClient
}

func NewNovelRepositoryImpl(ollamaClient *database.OllamaClient) port.NovelRepository {
	return &NovelRepositoryImpl{
		ollamaClient,
	}
}

func (r *NovelRepositoryImpl) GenerateNovel(content string) (*string, error) {
	// 送信モデル
	requestModel := model.OllamaRequestModel{
		Model:  "hf.co/mmnga/cyberagent-DeepSeek-R1-Distill-Qwen-14B-Japanese-gguf",
		System: r.getSystemPrompt(),
		Prompt: content,
	}
	// JSONを文字列化
	jsonData, err := json.Marshal(&requestModel)
	if err != nil {
		return nil, err
	}
	// リクエスト送信
	res, err := http.Post(
		r.ollamaClient.Endpoint,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// ボディ読み取り
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// JSON化
	responseModel := model.OllamaResponseModel{}
	if err := json.Unmarshal(body, &responseModel); err != nil {
		return nil, err
	}
	return &responseModel.Response, nil
}

// システムプロンプト
func (r *NovelRepositoryImpl) getSystemPrompt() string {
	return `
        あなたは小説家です
        夢の内容から小説を考えてください
    `
}
