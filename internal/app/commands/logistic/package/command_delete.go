package _package

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *PackageCommander) Delete(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")

	packageId, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		log.Println("failed to parse packageId", args)
		return
	}

	_, err = c.packageService.Remove(packageId)
	if err != nil {
		log.Printf("failed to delete package with ID = %d: %v", packageId, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("package with ID = %d have been successfully deleted", packageId),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", msg, err)
	}
}