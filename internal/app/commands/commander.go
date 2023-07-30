package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MaksimUlitin/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//ar rCommandes = map[string] func(c *Commander, msg tgbotapi.Message)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Offset int `json:"offset"`
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HendleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered form panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Parsed: %+v\n ", parsedData))

		c.bot.Send(msg)
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}

}
