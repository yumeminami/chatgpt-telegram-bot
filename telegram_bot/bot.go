package telegrambot

import (
	"chatgpt-telegram-bot/config"
	"log"

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

	// u := tgbotapi.NewUpdate(0)

	// updates := bot.GetUpdatesChan(u)
	// time.Sleep(5 * time.Second)
	// for len(updates) > 0 {
	// 	<-updates
	// }

	// for update := range updates {
	// 	if update.Message == nil {
	// 		continue
	// 	}

	// 	text := update.Message.Text
	// 	chat_id := update.Message.Chat.ID
	// 	// chat_username := update.Message.Chat.UserName

	// 	reply := Handler(&text)
	// 	if reply == nil {
	// 		continue
	// 	}
	// 	msg := tgbotapi.NewMessage(chat_id, *reply)
	// 	return_message, err := bot.Send(msg)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	log.Println(return_message)
	// }
	// select {}
}

func Handler(msg *string) *string {
	return msg
}

func RunBot() {

}

func StopBot() {

}
