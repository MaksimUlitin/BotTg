package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMassegText := " Here all the products \n\n"
	prodects := c.productService.List()
	for _, p := range prodects {
		outputMassegText += p.Title
		outputMassegText += "\n"

	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMassegText)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData))))
	c.bot.Send(msg)
}

