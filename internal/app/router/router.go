package router

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/serjleo/go_tg_bot/internal/service/product"
)

type CommandRouter struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommandRouter(bot *tgbotapi.BotAPI, productService *product.Service) *CommandRouter {
	return &CommandRouter{
		bot:            bot,
		productService: productService,
	}
}

func (r *CommandRouter) Process(update tgbotapi.Update) {
	switch update.Message.Text {
	case "/list":
		r.listCommand(update.Message, r.productService.List())
	case "/help":
		r.helpCommand(update.Message)
	default:
		r.defaultBehavior(update.Message)
	}
}

func (r *CommandRouter) helpCommand(msg *tgbotapi.Message) {
	r.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, `
	/help - help
	/list - list of products
`))
}

func (r *CommandRouter) listCommand(msg *tgbotapi.Message, products []product.Product) {
	message := "List of products \n\n"
	for i, p := range products {
		message += fmt.Sprintf("%d) %s \n", i+1, p.Title)
	}
	r.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, message))
}

func (r *CommandRouter) defaultBehavior(msg *tgbotapi.Message) {
	r.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, msg.Text))
}
