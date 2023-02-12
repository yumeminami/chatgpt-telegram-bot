package telegrambot

import (
	"chatgpt-telegram-bot/chatgpt"
	"chatgpt-telegram-bot/config"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitBot() {

	//get telegram bot token
	telegram_bot_token := config.GetConfig().Telegram_Bot_Token
	if telegram_bot_token == "" {
		panic("telegram bot token is empty")
	}

	// init telegram bot
	var err error
	bot, err = tgbotapi.NewBotAPI(telegram_bot_token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

}

func Handler(msg *string) *string {
	return msg
}

// func RunBot() {

// }

func StopBot() {

}

func RunBot(telegram_bot_token string) {
	bot, err := tgbotapi.NewBotAPI(telegram_bot_token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		reply := chatgpt.CreateCompletion(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

}
