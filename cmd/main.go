package main

import (
	"TelegrammBOTOPTIONS/botApi"
	"TelegrammBOTOPTIONS/store"
	"TelegrammBOTOPTIONS/trade"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var token string = "1122545961:AAGfVwD0Sowfqd_ICaBJG3n2CSSjBp1qs6o"

func main() {
	fmt.Println("BOT RUN")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/tgbot")
	if err != nil {
		fmt.Println(err)
	}
	s := store.NewStore(db)
	signalchan, stopsignal, static, testSignal := make(chan int, 1), make(chan int, 1), make(chan int, 1), make(chan int, 1) //id chat
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal("error run bot")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	go trade.ConnectBinary(signalchan, stopsignal, static, testSignal, bot, &s)
	botApi.RunBot(signalchan, stopsignal, static, testSignal, bot, &s)
}
