package _package

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *PackageCommander) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	outputMessage := tgbotapi.NewMessage(inputMsg.Chat.ID, "unknown command")

	_, err := c.bot.Send(outputMessage)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", outputMessage, err)
	}
}
