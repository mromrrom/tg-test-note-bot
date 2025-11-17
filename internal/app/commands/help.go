package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"help - help\n"+
			"list - список продуктов")
	msg.ReplyToMessageID = inputMsg.MessageID
	c.bot.Send(msg)
}

func init() {
	registredCommands["help"] = (*Commander).Help
}
