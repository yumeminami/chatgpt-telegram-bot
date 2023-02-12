package telegrambot

import (
	"chatgpt-telegram-bot/chatgpt"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func RunBot(telegram_bot_token string) {
	var err error
	bot, err = tgbotapi.NewBotAPI(telegram_bot_token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 3600

	updates := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	// hint the bot server

	for update := range updates {
		if update.Message == nil && update.EditedMessage == nil {
			continue
		}
		msg := Handler(&update)
		bot.Send(msg)
	}

}

func Handler(update *tgbotapi.Update) tgbotapi.MessageConfig {

	if update.Message != nil {
		if update.Message.IsCommand() {
			text := update.Message.Text[1:]
			commands, err := bot.GetMyCommands()
			if err != nil {
				return tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			}
			for _, command := range commands {
				fmt.Println(command.Command)
				if text == command.Command {
					return tgbotapi.NewMessage(update.Message.Chat.ID, command.Description)
				}
			}
		}
		reply := chatgpt.CreateCompletion(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// msg.ReplyToMessageID = update.Message.MessageID
		return msg

	}

	if update.EditedMessage != nil {
		relpy := chatgpt.CreateCompletion(update.EditedMessage.Text)
		msg := tgbotapi.NewMessage(update.EditedMessage.Chat.ID, relpy)
		msg.ReplyToMessageID = update.EditedMessage.MessageID
		return msg
	}

	return tgbotapi.NewMessage(update.Message.Chat.ID, "Error")
}
