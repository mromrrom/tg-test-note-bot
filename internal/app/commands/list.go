package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	var outputMsg string = "Here all the products:"
	for _, prod := range c.productService.List() {
		outputMsg += "\n" + prod.Tittle
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsg)
	//msg.ReplyToMessageID = inputMsg.MessageID

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func init() {
	registredCommands["list"] = (*Commander).List
}
