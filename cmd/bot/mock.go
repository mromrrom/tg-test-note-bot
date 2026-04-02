package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type tgMock struct {
}

func (c *tgMock) Send(cc tgbotapi.Chattable) (tgbotapi.Message, error) {
	fmt.Println(cc)
	return tgbotapi.Message{}, nil
}
