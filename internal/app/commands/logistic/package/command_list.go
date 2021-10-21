package _package

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *PackageCommander) List(inputMsg *tgbotapi.Message) {
	packages, err := c.packageService.List(0, Limit)

	if err != nil {
		log.Printf("failed to get list of packages: %v", err)
		return
	}

	var outputMessage strings.Builder

	outputMessage.WriteString("All packages: \n\n")
	for _, p := range packages {
		outputMessage.WriteString(fmt.Sprintf("%s\n", &p))
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMessage.String())

	if c.packageService.Count() > Limit {
		serializedData, err := json.Marshal(CallbackListData{
				Offset: Limit,
				Limit:  Limit,
			})

		if err != nil {
			log.Printf("failed to marshal CallbackListData: %v", err)
			return
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(NextPage, Command+string(serializedData)),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", outputMessage, err)
	}
}
