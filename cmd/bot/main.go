package main

import (
	"fmt"
	"github.com/AlexLuminare/demo-bot/internal/app/commands"
	"github.com/AlexLuminare/demo-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	godotenv.Load(".ENV")
	Token := os.Getenv("TELEGRAM_TOKEN")
	fmt.Println("TOKEN: ", Token)
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Panic(err)
	}

	//ВСЕ СЕРВИСЫ ИНИЦИАЛИЗИРУЕМ ЗДЕСЬ
	productService := product.NewService()
	commander := commands.NewCommandRouter(bot, productService)
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
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
