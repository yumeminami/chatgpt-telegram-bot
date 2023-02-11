package main

import (
	"chatgpt-telegram-bot/config"
	telegrambot "chatgpt-telegram-bot/telegram_bot"
)

func main() {
	config.InitConfig()
	telegrambot.InitBot()
	telegrambot.RunBot()
	telegrambot.StopBot()
}
