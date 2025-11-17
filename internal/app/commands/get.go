package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	//Строка с аргументами команды /Get
	args := inputMsg.CommandArguments()

	idx, ok := strconv.Atoi(args)
	if ok != nil {
		log.Println("Wrong args with /get command")
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("Fail to get product: inx=%d, err=%#v", idx, err)
		return
	}
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		product.Tittle,
	)

	//msg.ReplyToMessageID = inputMsg.MessageID
	_, err = c.bot.Send(msg)
	if err != nil {
		return
	}
}

func init() {
	registredCommands["get"] = (*Commander).Get
}
