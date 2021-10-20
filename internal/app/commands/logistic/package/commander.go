package _package

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
	"log"
)

type PackageService interface {
	Describe(packageId uint64) (*logistic.Package, error)
	List(cursor uint64, limit uint64) ([]logistic.Package, error)
	Create(logistic.Package) (uint64, error)
	Update(packageId uint64, _package logistic.Package) error
	Remove(packageId uint64) (bool, error)
	Count() uint64
}

type PackageCommander struct {
	bot            *tgbotapi.BotAPI
	packageService PackageService
}

func NewPackageCommander(bot *tgbotapi.BotAPI, service PackageService) *PackageCommander {
	return &PackageCommander{
		bot:              bot,
		packageService: service,
	}
}

func (c *PackageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *PackageCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}