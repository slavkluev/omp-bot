package _package

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *PackageCommander) Get(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	packageId, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		log.Println("failed to parse packageId", args)
		return
	}

	_package, err := c.packageService.Describe(packageId)
	if err != nil {
		log.Printf("failed to find package with ID = %d: %v", packageId, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprint(_package),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", msg, err)
	}
}