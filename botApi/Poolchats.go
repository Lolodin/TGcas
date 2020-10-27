package botApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

type PoolChats struct {
	Pool map[int]int
	Bot *tgbotapi.BotAPI
	Timer int
	Signal chan struct{}
}

func(p *PoolChats) SendMessage(msg string) {
	for _, v := range p.Pool {
		msg:= tgbotapi.NewMessage(int64(v), msg)
		p.Bot.Send(msg)
	}
}
func (p *PoolChats) AddChat(id int) {
	p.Pool[id] = id
}
func (p *PoolChats) OffChat(id int) {
	delete(p.Pool, id)
}
func NewPool(Bot *tgbotapi.BotAPI, timer int) PoolChats {
	p := PoolChats{}
	p.Bot = Bot
	p.Timer = timer
	p.Pool = make(map[int]int, 10)
	p.Signal = make(chan struct{}, 1)
	go func() {
		for {
			time.Sleep(time.Duration(p.Timer) * time.Second)
			p.Signal<- struct{}{}
		}
	}()
	return  p
}
