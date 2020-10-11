package main

import (
	"TelegrammBOTOPTIONS/botApi"
	"TelegrammBOTOPTIONS/trade"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)
var token string= "1122545961:AAGfVwD0Sowfqd_ICaBJG3n2CSSjBp1qs6o"
func main() {
	fmt.Println("BOT RUN")
	signalchan := make(chan int, 1) //id chat
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("error run bot")
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	go trade.ConnectBinary(signalchan, bot)
	botApi.RunBot(signalchan, bot)
}
