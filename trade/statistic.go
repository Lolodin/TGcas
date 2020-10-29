package trade

import (
	"TelegrammBOTOPTIONS/botApi"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type statistic struct {
	Up int
	Down int
	Winrate float64
	Bot *tgbotapi.BotAPI
}
func(s *statistic) GetStatistic() {
	global:= float64(s.Up+s.Down)
	s.Winrate = (float64(s.Up)/global)*100
  str:=fmt.Sprint("Верно: ", s.Up, " Ошибок: ",s.Down, " Винрейт ", s.Winrate)
  msg := tgbotapi.NewMessage(botApi.ADMINCHAT, str)
  s.Bot.Send(msg)
}
func(s *statistic) WrongAdd() {
	s.Down++
}
func(s *statistic) GoodAdd() {
	s.Up++
}
