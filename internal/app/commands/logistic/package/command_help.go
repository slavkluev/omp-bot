package _package

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *PackageCommander) Help(inputMsg *tgbotapi.Message) {
	helpMessage := `supported commands:
/help__logistic__package - help
/get__logistic__package {package_id} - get package
/list__logistic__package {start_package_id} {limit} {page_limit} - list packages
/delete__logistic__package {package_id} - delete package
/new__logistic__package {weight} {width} {height} {length} - create new package
/edit__logistic__package {package_id} {weight} {width} {height} {length} - edit package`

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, helpMessage)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", msg, err)
	}
}
