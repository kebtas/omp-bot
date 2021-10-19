package verification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/security/verification"
)

type SecurityVerificationCommander struct {
	bot                 *tgbotapi.BotAPI
	verificationService *verification.VerificationService
}

func NewSecurityVerificationCommander(
	bot *tgbotapi.BotAPI,
) *SecurityVerificationCommander {
	verificationService := verification.NewVerificationService([]verification.Verification{
		{Title: "one"},
		{Title: "two"},
		{Title: "three"},
		{Title: "four"},
		{Title: "five"},
	})

	return &SecurityVerificationCommander{
		bot:                 bot,
		verificationService: verificationService,
	}
}

func (c *SecurityVerificationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	default:
		log.Printf("SecurityVerificationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *SecurityVerificationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	}
}