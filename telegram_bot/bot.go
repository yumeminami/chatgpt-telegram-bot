package telegrambot

import (
	"chatgpt-telegram-bot/chatgpt"
	"errors"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var msgMap map[int]string

const (
	MainMenuText = "*Main Menu*\n" +
		"*/completions* - _Creates a completion for the provided prompt and parameters_\n" +
		"*/edit* - _Creates a new edit for the provided input, instruction, and parameters._\n" +
		"*/images* - _Creates an image given a prompt._\n" +
		"*/images_edit* - _Creates an edited or extended image given an original image and a prompt._\n"
)

// make a command usage map
var commandUsage = map[string]string{
	"completions": "Enter a prompt and I will complete it for you.",
	"edit":        "You can reply to message I have sent you and I will regenerate the message based the previous one.",
	"images":      "Enter a prompt and I will create a photo.",
	"images_edit": "Reply the image message and I will edited or extended image.",
}

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
		if update.InlineQuery != nil {
		}
		msg := Handler(&update)
		bot.Send(msg)
	}

}

func Handler(update *tgbotapi.Update) tgbotapi.MessageConfig {
	if update.Message != nil || update.EditedMessage != nil {
		var message *tgbotapi.Message
		if update.Message != nil {
			message = update.Message
		} else {
			message = update.EditedMessage
		}
		action := tgbotapi.NewChatAction(message.Chat.ID, "typing")
		bot.Send(action)
		// command
		if message.IsCommand() {
			log.Printf("Command: [%s] %s", message.From.UserName, message.Command())
			text := message.Text[1:]
			if text == "start" {
				return MainMenu(message.Chat.ID)
			}
			reply, err := HandlerCommand(text)
			if err != nil {
				return tgbotapi.NewMessage(message.Chat.ID, err.Error())
			}
			return tgbotapi.NewMessage(message.Chat.ID, reply)

		}
		// edit
		if message.ReplyToMessage != nil {
			log.Printf("Edit: [%s] %s", message.From.UserName, message.Text)
			if text, ok := msgMap[message.ReplyToMessage.MessageID-1]; ok {
				reply, err := chatgpt.CreateEdit(text, message.Text)
				if err != nil {
					return tgbotapi.NewMessage(message.Chat.ID, err.Error())
				}
				msg := tgbotapi.NewMessage(message.Chat.ID, reply)
				msgMap[message.MessageID] = reply
				return msg
			}
		}
		// completion
		log.Printf("Completion: [%s] %s", message.From.UserName, message.Text)
		reply := chatgpt.CreateCompletion(message.Text)
		msg := tgbotapi.NewMessage(message.Chat.ID, reply)
		msgMap[message.MessageID] = reply
		return msg

	}

	if update.CallbackQuery != nil {
		log.Printf("CallbackQuery: [%s] %s", update.CallbackQuery.From.UserName, update.CallbackQuery.Data)
		replay, err := HandlerCommand(update.CallbackQuery.Data)
		if err != nil {
			return tgbotapi.NewMessage(update.CallbackQuery.From.ID, err.Error())
		}
		return tgbotapi.NewMessage(update.CallbackQuery.From.ID, replay)
	}

	return tgbotapi.NewMessage(0, "Error")
}

func MainMenu(chatID int64) tgbotapi.MessageConfig {
	text := MainMenuText
	completions_btn := tgbotapi.NewInlineKeyboardButtonData("Completions", "completions")
	edit_btn := tgbotapi.NewInlineKeyboardButtonData("Edit", "edit")
	images_btn := tgbotapi.NewInlineKeyboardButtonData("Images", "images")
	images_edit_btn := tgbotapi.NewInlineKeyboardButtonData("Images Edit", "images_edit")
	row_one := tgbotapi.NewInlineKeyboardRow(completions_btn, edit_btn)
	row_two := tgbotapi.NewInlineKeyboardRow(images_btn, images_edit_btn)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row_one, row_two)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "Markdown"
	return msg
}

func HandlerCommand(text string) (string, error) {

	commands, err := bot.GetMyCommands()
	if err != nil {
		return "", err
	}
	for _, command := range commands {
		if text == command.Command {
			return commandUsage[command.Command], nil
		}
	}
	return "", errors.New("Command not found")
}
