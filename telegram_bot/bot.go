package telegrambot

import (
	"chatgpt-telegram-bot/chatgpt"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var msgMap map[int]string

const (
	MainMenuText = "*Main Menu*\n" +
		"*Completion* - _Just enter the prompt_\n" +
		"*Edit* - _Reply the to message as input with message as instruction_\n"
)

func RunBot(telegram_bot_token string) {
	var err error
	bot, err = tgbotapi.NewBotAPI(telegram_bot_token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 600

	updates := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	msgMap = make(map[int]string)

	for update := range updates {
		msg := Handler(&update)
		bot.Send(msg)
	}

}

func Handler(update *tgbotapi.Update) tgbotapi.MessageConfig {
	if update.Message != nil {
		action := tgbotapi.NewChatAction(update.Message.Chat.ID, "typing")
		bot.Send(action)
		if update.Message.IsCommand() {
			text := update.Message.Text[1:]
			if text == "start" {
				return MainMenu(update.Message.Chat.ID)
			}

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
		if update.Message.ReplyToMessage != nil {
			fmt.Println("OK")
			if text, ok := msgMap[update.Message.ReplyToMessage.MessageID-1]; ok {
				reply, err := chatgpt.CreateEdit(text, update.Message.Text)
				if err != nil {
					return tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				msgMap[update.Message.MessageID] = reply
				return msg
			}
		}
		fmt.Println(update.Message.Text)
		reply := chatgpt.CreateCompletion(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msgMap[update.Message.MessageID] = reply
		return msg

	}

	if update.EditedMessage != nil {
		action := tgbotapi.NewChatAction(update.EditedMessage.Chat.ID, "typing")
		bot.Send(action)
		relpy := chatgpt.CreateCompletion(update.EditedMessage.Text)
		msg := tgbotapi.NewMessage(update.EditedMessage.Chat.ID, relpy)
		msg.ReplyToMessageID = update.EditedMessage.MessageID
		return msg
	}

	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "completion":
			return tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Enter the prompt")
		case "edit":
			return tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Reply the Message")
		}
	}

	return tgbotapi.NewMessage(0, "Error")
}

func MainMenu(chatID int64) tgbotapi.MessageConfig {
	text := MainMenuText
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Completion", "completion")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Edit", "edit")
	row := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "Markdown"
	return msg
}
