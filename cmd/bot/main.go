package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlexLuminare/demo-bot/internal/app/commands"
	"github.com/AlexLuminare/demo-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.ENV")
	Token := os.Getenv("TELEGRAM_TOKEN")
	fmt.Println("TOKEN: ", Token)
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Panic(err)
	}

	//var m tgMock
	//ВСЕ СЕРВИСЫ ИНИЦИАЛИЗИРУЕМ ЗДЕСЬ
	productService := product.NewService()
	commander := commands.NewCommandRouter(bot, productService)
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u) // 26.03.2026 замокать updates на гитхабе это было
	// https://github.com/bot-api/telegram/blob/master/testutils/mocks.go
	// ищем способы замокать телегу
	if err != nil {
		log.Panic(err)
	}

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()
	//var msg tgbotapi.MessageConfig

	for update := range updates {
		commander.HandleUpdate(&update)
		//msg.ReplyToMessageID = update.Message.MessageID
	}
}

// заняться той частью где надо в нужную дату отправлять сообщения
// 1 - вычитать весь файл и найти нужные даты (можно для удобства подчистить файл),
// сохранять в структуру напр все задания на текущий день, неправильные даты можно пропускать
// 2 -
// 3 -
//
//
