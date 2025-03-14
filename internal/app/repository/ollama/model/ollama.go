package model

type OllamaRequestModel struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	System string `json:"system"`
}

type OllamaResponseModel struct {
	Response string `json:"response"`
}
