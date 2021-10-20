package logistic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	packageCommander "github.com/ozonmp/omp-bot/internal/app/commands/logistic/package"
	"github.com/ozonmp/omp-bot/internal/app/path"
	packageService "github.com/ozonmp/omp-bot/internal/service/logistic/package"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LogisticCommander struct {
	bot              *tgbotapi.BotAPI
	packageCommander Commander
}

func NewLogisticCommander(
	bot *tgbotapi.BotAPI,
) *LogisticCommander {
	return &LogisticCommander{
		bot:              bot,
		packageCommander: packageCommander.NewPackageCommander(bot, packageService.NewPackageService()),
	}
}

func (c *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "package":
		c.packageCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "package":
		c.packageCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
