package _package

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
	"log"
	"strconv"
	"strings"
)

func (c *PackageCommander) Edit(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	args := inputMsg.CommandArguments()
	args = strings.Trim(args, " ")
	argsArray := strings.Split(args, " ")

	if len(argsArray) != 5 {
		log.Printf("need to be 5 args")
		return
	}

	packageId, err := strconv.ParseUint(argsArray[0], 10, 64)
	if err != nil {
		log.Printf("failed to parse packageId: %v", err)
		return
	}

	weight, err := strconv.ParseUint(argsArray[1], 10, 64)
	if err != nil {
		log.Printf("failed to parse weight: %v", err)
		return
	}

	width, err := strconv.ParseUint(argsArray[2], 10, 64)
	if err != nil {
		log.Printf("failed to parse width: %v", err)
		return
	}

	height, err := strconv.ParseUint(argsArray[3], 10, 64)
	if err != nil {
		log.Printf("failed to parse height: %v", err)
		return
	}

	length, err := strconv.ParseUint(argsArray[4], 10, 64)
	if err != nil {
		log.Printf("failed to parse length: %v", err)
		return
	}

	_package := logistic.NewPackage(weight, width, height, length)

	err = c.packageService.Update(packageId, _package)
	if err != nil {
		log.Printf("failed to edit package with ID = %d: %v", packageId, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("package with ID = %d have been successfully edited", packageId),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("error sending reply message (%v) to chat: %v", msg, err)
	}
}