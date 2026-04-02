package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// type Tasks struct {
//	Task[] string
//}

// сегодня, завтра, послезавтра, выходные, следующие выходные,
// понедельник, вторник среда четверг пятница суббота воскресенье, ...

func (c *Commander) SendDailyMessage(inputMsg *tgbotapi.Message) {

}

// 13.02.26 -> продолжить написание этой части
func (c *Commander) SendTodayMessages(inputMsg *tgbotapi.Message) {
	tasks := findTodayTasks(inputMsg.Chat.ID)
	tasks = append([]string{"Сегодняшние задачи: "}, tasks...)
	taskString := strings.Join(tasks, "\n")
	resp := tgbotapi.NewMessage(inputMsg.Chat.ID, taskString)
	c.bot.Send(resp)
}

// находит в файле все записи с текущей датой и возвращает список этих записей
// 13.02.26 -> делаю это
func findTodayTasks(chatId int64) []string {

	var Tasks []string

	file, err := os.Open("my_log.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Создаем сканер
	scanner := bufio.NewScanner(file)

	// Читаем построчно
	for scanner.Scan() {
		line := scanner.Text() // Получаем строку
		splitedLine := strings.Split(line, "; ")
		if len(splitedLine) > 2 {

			if splitedLine[1] != strconv.FormatInt(chatId, 10) {
				continue
			}

			intTimestamp, err := strconv.ParseInt(splitedLine[2], 10, 64)
			if err != nil {
				fmt.Println("Error parsing Int64 from string")
				Tasks = append(Tasks, "Error parsing Int64 from string")
				return Tasks
			}

			DeadlineDate := time.Unix(intTimestamp, 0)

			if DeadlineDate.YearDay() == time.Now().YearDay() && DeadlineDate.Year() == time.Now().Year() { // TODO: доделать сравнение, чтобы оно было корректным
				Tasks = append(Tasks, splitedLine[3])
			}
		}
	}

	// Проверяем ошибки сканирования
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Tasks
}

func init() {
	registredCommands["getTasks"] = (*Commander).SendTodayMessages
}

//TODO: функция сохранения чтобы он сохранял и текст и chat id и дату
