package main

import (
	"chatgpt-telegram-bot/chatgpt"
	"chatgpt-telegram-bot/config"
	telegrambot "chatgpt-telegram-bot/telegram_bot"
)

func main() {
	config.InitConfig()
	chatgpt.InitClient()
	telegrambot.RunBot(config.GetConfig().Telegram_Bot_Token)
}

// package main

// import (
// 	"context"
// 	"fmt"

// 	gogpt "github.com/sashabaranov/go-gpt3"
// )

// func main() {
// 	c := gogpt.NewClient("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")
// 	ctx := context.Background()

// 	req := gogpt.CompletionRequest{
// 		Model:     gogpt.GPT3Ada,
// 		MaxTokens: 1000,
// 		Prompt:    "Lorem ipsum",
// 	}
// 	resp, err := c.CreateCompletion(ctx, req)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(resp.Choices[0].Text)
// }
