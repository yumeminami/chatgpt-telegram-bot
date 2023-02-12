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
