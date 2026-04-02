package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) SaveToFile(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	file, err := os.OpenFile("my_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Всегда закрывайте файл
	msgText := inputMsg.Text
	msgText = strings.Replace(msgText, "/save ", "", 1)

	var deadlineDateInTimeFormat time.Time

	// дз убрать команду save из сообщения, и сохранять дату из сообщения а не ту когда сообщение было отправлено
	// если будет время то придумать каким образом задать дату
	candidate, err := WordDatetimeSearcher(inputMsg.Text)
	//TODO: 13.03.2026 ДЗ доописать функцию WordDatetimeSearcher и протестировать что получилось, если нужно SaveToFile поправить
	if err != nil {
		fmt.Println("error", err)

		candidate2, err2 := DateTimeSearcher(inputMsg.Text)
		if err2 != nil {
			deadlineDateInTimeFormat = time.Now()
			fmt.Println("error", err2)
		}
		if err2 == nil {
			deadlineDateInTimeFormat = candidate2
		}
	} // 19.03.2026 дз дописать функцию DateTimeSearcher чтобы она работала правильно
	// 19.03.2026 дз дописать функцию WordDatetimeSearcher
	if err == nil {
		deadlineDateInTimeFormat = candidate
	}

	data := []byte(fmt.Sprintf("%v; %v; %v; %s\n", inputMsg.Date, inputMsg.Chat.ID, deadlineDateInTimeFormat.Unix(), msgText))
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	resp := tgbotapi.NewMessage(inputMsg.Chat.ID, "Saved")
	c.bot.Send(resp)
}

var KeyDays = []string{
	"сегодня",
	"завтра",
	"послезавтра",
	"выходные",
	//"следующие выходные",
	"понедельник",
	"вторник",
	"сред",
	"четверг",
	"пятница",
	"суббота",
	"воскресенье",
}

// добавить поиск по ключевым словам: сегодня, завтра, на следующей неделе, в следующий понедельник, итп
// сегодня, завтра, послезавтра, выходные, следующие выходные,
// понедельник, вторник среда четверг пятница суббота воскресенье, ...
// возможно хранить тип datetime
// next: посмотреть поддерживает ли библиотека обработку сообщений БЕЗ ключевых слов (команд типа /save )
// TODO: ДЗ 13.03 посмотреть пакет со строками strings, изучить

func WordDatetimeSearcher(inputMsg string) (time.Time, error) {
	for _, v := range KeyDays {
		if strings.Contains(inputMsg, v) && v == "сегодня" {
			return time.Now(), nil
		}
		if strings.Contains(inputMsg, v) && v == "завтра" {
			return time.Now().AddDate(0, 0, 1), nil
		}
		if strings.Contains(inputMsg, v) && v == "послезавтра" {
			return time.Now().AddDate(0, 0, 2), nil
		}
		if strings.Contains(inputMsg, v) && v == "выходные" {
			weekDay := time.Now().Weekday()
			if weekDay > 0 && weekDay < 6 {
				deltaTime := 6 - weekDay
				return time.Now().AddDate(0, 0, int(deltaTime)), nil
			}
			if weekDay == 0 {
				return time.Now().AddDate(0, 0, 6), nil
			}
			if weekDay == 6 {
				return time.Now().AddDate(0, 0, 7), nil
			}
		}
		if strings.Contains(inputMsg, v) && v == "понедельник" {
			weekDay := time.Now().Weekday()
			if weekDay == 0 {
				return time.Now().AddDate(0, 0, 1), nil
			}
			if weekDay == 1 {
				return time.Now(), nil
			}
			if weekDay >= 2 && weekDay <= 6 {
				delta := 8 - weekDay
				return time.Now().AddDate(0, 0, int(delta)), nil
			}
		}
	}
	return time.Time{}, errors.New("no word date found")
}

func DateTimeSearcher(inputMsg string) (time.Time, error) {
	re := regexp.MustCompile(`\d{1,2}\.\d{1,2}\.\d{2,4}`)
	date := re.Find([]byte(inputMsg))
	if date == nil {
		err := errors.New("no date found")
		return time.Time{}, err
	}
	deadlineDateInTimeFormat, err := time.Parse("02.01.2006", string(date))
	if err != nil {
		fmt.Println("error", err)
	}
	return deadlineDateInTimeFormat, nil
}

func init() {
	registredCommands["save"] = (*Commander).SaveToFile
}

// error(*time.ParseError) *{Layout: "2006-01-02", Value: "15.02.26", LayoutElem: "2006", ValueElem: "15.02.26", Message: ""}
// fmt.Println("error", err)
