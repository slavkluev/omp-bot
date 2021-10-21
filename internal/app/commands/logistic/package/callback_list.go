package _package

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strings"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

const (
	Command      = "logistic__package__list__"
	Limit        = 5
	PreviousPage = "<"
	NextPage     = ">"
)

func (c *PackageCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	log.Printf("[%s] %s", callback.From.UserName, callbackPath.CallbackData)

	callbackListData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &callbackListData)

	if err != nil {
		log.Printf("Failed to parse data %v", callbackPath.CallbackData)
		return
	}

	packages, err := c.packageService.List(callbackListData.Offset, callbackListData.Limit)

	if err != nil {
		log.Printf("failed to get list of packages: %v", err)
		return
	}

	var outputMessage strings.Builder

	for _, p := range packages {
		outputMessage.WriteString(fmt.Sprintf("%s\n", &p))
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMessage.String())

	buttons, err := createButtons(callbackListData, c.packageService.Count())

	if err != nil {
		log.Printf("failed to marshal CallbackListData: %v", err)
		return
	}

	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", msg, err)
	}
}

func createButtons(parsedData CallbackListData, packagesCount uint64) ([]tgbotapi.InlineKeyboardButton, error) {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0, 2)

	if parsedData.Offset > 0 {
		previousPageStart := parsedData.Offset - Limit

		if previousPageStart < 0 {
			previousPageStart = 0
		}

		serializedData, err := json.Marshal(CallbackListData{
			Offset: previousPageStart,
			Limit:  Limit,
		})

		if err != nil {
			return []tgbotapi.InlineKeyboardButton{}, err
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(PreviousPage, Command+string(serializedData)))
	}

	nextPageCursor := parsedData.Offset + parsedData.Limit

	if packagesCount > nextPageCursor {
		serializedData, err := json.Marshal(CallbackListData{
			Offset: nextPageCursor,
			Limit:  Limit,
		})

		if err != nil {
			return []tgbotapi.InlineKeyboardButton{}, err
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(NextPage, Command+string(serializedData)))
	}

	return buttons, nil
}
