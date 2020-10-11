package botApi

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

const (
	START = "start"
	TEST  = "test"
	GETTEST = "gettest"
)
type Signal struct {
	IDchat int64 // id chat from Bot
}



func RunBot(signal chan int, bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		fmt.Println(update.Message.Command())
		//fmt.Println(update.Message.IsCommand())
		//fmt.Println(update.Message.Text)
		// Add logic here
		switch update.Message.Command() {
		case START:
			fmt.Println(update.Message.From.ID)
			msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "Вас приветсвует система сигналов опционов, для получения тестового доступа, введите /gettest")
			bot.Send(msg)
		case TEST:
			signal <- update.Message.From.ID
		case GETTEST:
			userID := update.Message.From.ID
			Time:= time.Now().Unix()

		default:
			msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "Команда не валидна")
			bot.Send(msg)

		}
	}

}