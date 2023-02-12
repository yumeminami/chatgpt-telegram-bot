package chatgpt

import (
	"chatgpt-telegram-bot/config"
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var client *gogpt.Client
var ctx context.Context

func InitClient() {
	client = gogpt.NewClient(config.GetConfig().Openai_API_Key)
	ctx = context.Background()
}

func CreateCompletion(msg string) string {
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   4000,
		Prompt:      msg,
		Temperature: 0.7,
		TopP:        1,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		return ""
	}
	fmt.Println(resp.Choices[0].Text)
	return resp.Choices[0].Text
}
