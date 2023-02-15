package chatgpt

import (
	"chatgpt-telegram-bot/config"
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var client *gogpt.Client
var ctx context.Context

func InitClient() {
	client = gogpt.NewClient(config.GetConfig().Openai_API_Key)
	ctx = context.Background()
}
