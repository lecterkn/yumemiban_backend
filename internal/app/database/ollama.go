package database

import "os"

type OllamaClient struct {
	Endpoint string
}

func GetOllamaClient() *OllamaClient {
	return &OllamaClient{
		Endpoint: getEndpointUrl(),
	}
}

func getEndpointUrl() string {
	endpoint, ok := os.LookupEnv("OLLAMA_ENDPOINT_URL")
	if !ok {
		panic("環境変数\"OLLAMA_ENDPOINT_URL_\"が設定されていません")
	}
	return endpoint
}
