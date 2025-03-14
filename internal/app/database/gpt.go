package database

import (
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetChatGPTClient() *openai.Client {
	return openai.NewClient(getOpenAPIToken())
}

func getOpenAPIToken() string {
	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		panic("環境変数に\"OPENAI_API_LEY\"が設定されていません")
	}
	return apiKey
}
