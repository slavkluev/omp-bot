package _package

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
	"log"
	"strconv"
	"strings"
)

func (c *PackageCommander) New(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")
	argsArray := strings.Split(args, " ")

	if len(argsArray) != 4 {
		log.Printf("need to be 4 args")
		return
	}

	weight, err := strconv.ParseUint(argsArray[0], 10, 64)
	if err != nil {
		log.Printf("failed to parse weight: %v", err)
		return
	}

	width, err := strconv.ParseUint(argsArray[1], 10, 64)
	if err != nil {
		log.Printf("failed to parse width: %v", err)
		return
	}

	height, err := strconv.ParseUint(argsArray[2], 10, 64)
	if err != nil {
		log.Printf("failed to parse height: %v", err)
		return
	}

	length, err := strconv.ParseUint(argsArray[3], 10, 64)
	if err != nil {
		log.Printf("failed to parse length: %v", err)
		return
	}

	newPackage := logistic.NewPackage(weight, width, height, length)
	packageId, err := c.packageService.Create(newPackage)

	if err != nil {
		log.Printf("failed to create package: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("created package with ID = %d", packageId),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", msg, err)
	}
}
