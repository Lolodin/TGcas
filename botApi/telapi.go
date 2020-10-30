package botApi

import (
	"TelegrammBOTOPTIONS/store"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	START     = "/start"
	GETSIG    = "Включить бота"
	OFFSIG    = "Выключить бота"
	PAY       = "Оплата"
	FREE      = "Бесплатно"
	TEXTPAY   = "Выберите длительность подписки"
	ADMINCHAT = -377657292
	GETLINK   = "GetLink"
	GETSTATIC   = "getStatic"
	GETTEST   = "Тестовый доступ"
	SEND		= "sendall"
	PATTERN   = `^\w+@\w+\.\w+$`
)
type Signal struct {
	IDchat int64 // id chat from Bot
}



func RunBot(signal, stopsignal,static, testSignal chan int, bot *tgbotapi.BotAPI, stor *store.MySQL) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "1Mount":
				fmt.Println("1 Mount select")
				msg:= tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Стоимость участия на *1 месяц* составляет *10 000 рублей*\n\n" +
					"Оплатить можете картой нажав на кнопку ниже")
				numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Оплатить 1 месяц", "https://money.yandex.ru/to/410017694197716"),
					),
				)
				msg.ReplyMarkup = numericKeyboard
				msg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg)
				msg2:=tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "После оплаты отправьте скриншот транзакции в этот чат")
				bot.Send(msg2)

			case "2Mount":
				msg:= tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Стоимость участия на *2 месяц* составляет *15 000 рублей*\n\n" +
					"Оплатить можете картой нажав на кнопку ниже")
				numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Оплатить 2 месяца", "https://money.yandex.ru/to/410017694197716"),
					),
				)
				msg.ReplyMarkup = numericKeyboard
				msg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg)
				msg2:=tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "После оплаты отправьте скриншот транзакции в этот чат")
				bot.Send(msg2)
			case "3Mount":
				msg:= tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Стоимость участия на *3 месяц* составляет *20 000 рублей*\n\n" +
					"Оплатить можете картой нажав на кнопку ниже")
				numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Оплатить 3 месяца", "https://money.yandex.ru/to/410017694197716"),
					),
				)
				msg.ReplyMarkup = numericKeyboard
				msg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg)
				msg2:=tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "После оплаты отправьте скриншот транзакции в этот чат")
				bot.Send(msg2)
			case GETLINK:
				msg2:=tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "https://record.binary.com/_QxtIBHlCm1lM5vlemjwB2mNd7ZgqdRLk/1/\n\n" +
					"После Регистрации пришлите ниже адрес вашей электронной почты.")
				bot.Send(msg2)
			default:
				fmt.Println(update.InlineQuery.Query, "test")

			}
		}
		if update.Message == nil {
			continue
		}
		if update.Message.Chat.ID == ADMINCHAT {

			args := update.Message.CommandArguments()
			command:= update.Message.Command()
			if command == SEND {
				u:=stor.GetUserList()
				for _, v := range u.List {
					msg:=tgbotapi.NewMessage(int64(v), args)
					bot.Send(msg)
				}
			}
			if command == GETSTATIC {
				static <- 1
			}
			arr:=strings.Fields(args)
			if len(arr)<2 {
				continue
			}

			if command == "yes" {
				userID := arr[0]
				timeM  := arr[1]
				i, err :=strconv.Atoi(userID)
				if err != nil {
					msg:=tgbotapi.NewMessage(ADMINCHAT, "Неверные аргументы")
					bot.Send(msg)
				}
				m, err :=strconv.Atoi(timeM)
				if err != nil {
					msg:=tgbotapi.NewMessage(ADMINCHAT, "Неверные аргументы")
					bot.Send(msg)
				}
				err = stor.AddSubscription(i,m)
				if err != nil {
					msg:=tgbotapi.NewMessage(ADMINCHAT, "Ошибка добавления пользователя")
					bot.Send(msg)
				}

				msg:=tgbotapi.NewMessage(ADMINCHAT, "Юзер добавлен в подписчики")
				usermsg:= tgbotapi.NewMessage(int64(i), "Вам открыт доступ к сигналам для подключения включите бота в меню кнопкой `" +GETSIG+"`")
				bot.Send(msg)
				menu:= tgbotapi.ReplyKeyboardMarkup{}
				button1:= tgbotapi.KeyboardButton{}

				var row2 []tgbotapi.KeyboardButton
				button1.Text = "Включить бота"
				row2 = append(row2,button1)
				button1.Text = "Выключить бота"
				row2 = append(row2,button1)
				menu.Keyboard = append(menu.Keyboard, row2)
				menu.ResizeKeyboard = true
				menu.Selective = false





				usermsg.ReplyMarkup = menu
				bot.Send(usermsg)
			} else {
				userID := arr[0]
				i, err :=strconv.Atoi(userID)
				if err != nil {
					msg:=tgbotapi.NewMessage(ADMINCHAT, "Неверные аргументы")
					bot.Send(msg)
				}
				usermsg:= tgbotapi.NewMessage(int64(i), "Вам отказали в доступе к сигналам")
				bot.Send(usermsg)

			}



		}



		switch update.Message.Text {
		case START:
			fmt.Println(update.Message.From.ID)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "*Закрытый бот для трейдеров бо - информация о боте ↓*\n\n" +
				"*Уже 5-ый год* мы активно занимаемся трейдингом и заработали на бинарных опционах по-настоящему большую прибыль, поэтому  считаем, что имеем право и знания, которыми будет полезно делиться с вами.\n\n" +
				"Торговля ведется на бинарных опционах у брокера Binary.com. На данный момент это самый честный брокер, который сто процентно выплачивает прибыль. Других таких нет.\n\n" +
				"Торговля ведется на торговой паре EUR/USD, торговля не тиковая, что придает боту надежность и приемущество перед тиковыми ботами, где тики зависят от пинга, времени, и зависаний, здесь такого нет.\n\n" +
				"*Доступ к боту предоставляется ограниченному числу лиц.*\n\n" +
				"Если у вас нет возможности преобрести доступ к боту, то наш сервис предлагает альтернативный бесплатный вариант, который ориентирован на долгосрочное сотрудничество в соответствии с философией win - win.\n\n" +
				"*Что такое философия win - win?*\n\n" +
				"Это взаимодействие, от которого все участники могут получить прибыль : Вы можете получить доступ к боту бесплатно, зарегистрировашись по нашей реферальной ссылке на binary.com.\n\n" +
				"Таким образом вы будете получать наши сигналы, взамен на то, что мы будем получать комиссионые от вашего *торгового объема* у брокера.\n\n" +
				"*Что сейчас есть в закрытом боте для трейдеров?*\n\n▴ Понятные сигналы : торговая пара, точка входа, время эксперации.\n▴ " +
				"Авторская аналитика, которая публикуется только для членов бота.\n▴ " +
				"Прямое общение с администрацией.\n\n" )
			msg.ParseMode = "markdown"


			menu:= tgbotapi.ReplyKeyboardMarkup{}
			var row []tgbotapi.KeyboardButton
			button1:= tgbotapi.KeyboardButton{}
			button1.Text = "Оплата"


			row = append(row,button1)
			button1.Text = "Бесплатно"
			row = append(row,button1)
			button1.Text = "Контакты"
			row = append(row,button1)
			var row2 []tgbotapi.KeyboardButton
			button1.Text = "Включить бота"
			row2 = append(row2,button1)
			button1.Text = "Выключить бота"
			row2 = append(row2,button1)
			var row3 []tgbotapi.KeyboardButton
			button1.Text = "Тестовый доступ"
			row3 = append(row3,button1)

			menu.Keyboard = append(menu.Keyboard, row)
			menu.Keyboard = append(menu.Keyboard, row2)
			menu.Keyboard = append(menu.Keyboard, row3)
			menu.ResizeKeyboard = true
			menu.Selective = false





			msg.ReplyMarkup = menu
			bot.Send(msg)








		case "Контакты":
			msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "По всем вопросам пишите @Buffettadmin")
			bot.Send(msg)

		case GETSIG:
			user := update.Message.From.ID
			u, err :=stor.GetUserByID(user)
			if err != nil {
				fmt.Println(err)
				msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "У вас нет доступа к сигналам")
				bot.Send(msg)
				continue
			}
			arr:=strings.Split(u.Subscription, "-")
			var sdate []int
			sdate = make([]int, 3,3)
			for i, v := range arr {
				sdate[i], err = strconv.Atoi(v)
				if err != nil {
					fmt.Println(err)
				}
			}
			 t:=time.Date(sdate[0],time.Month(sdate[1]),sdate[2],0,0,0,0,time.UTC )
			 t2:=t.AddDate(0,u.TimeSub, 0)
			 b:=t.After(t2)
			if u == nil || b {
				msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "У вас нет доступа к сигналам")
				bot.Send(msg)
				continue
			}
			msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "Подключаем к сигналам, для остановки сигналов выберите в меню кнопку `"+OFFSIG+"`")
			menu:= tgbotapi.ReplyKeyboardMarkup{}
			button1:= tgbotapi.KeyboardButton{}

			var row2 []tgbotapi.KeyboardButton
			button1.Text = "Включить бота"
			row2 = append(row2,button1)
			button1.Text = "Выключить бота"
			row2 = append(row2,button1)

			var row3 []tgbotapi.KeyboardButton
			button1.Text = "Тестовый доступ"
			row3 = append(row3,button1)
			menu.Keyboard = append(menu.Keyboard, row2)
			menu.Keyboard = append(menu.Keyboard, row3)
			menu.ResizeKeyboard = true
			menu.Selective = false





			msg.ReplyMarkup = menu

			bot.Send(msg)

			signal <- int(update.Message.Chat.ID)
		case OFFSIG:
		stopsignal <-int(update.Message.Chat.ID)
		case GETTEST:
			testSignal<-int(update.Message.Chat.ID)

		case PAY:
			msg:=tgbotapi.NewMessage(update.Message.Chat.ID, TEXTPAY)
			numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Cтать участником клуба на 1 месяц", "1Mount"),
				),
				tgbotapi.NewInlineKeyboardRow(

					tgbotapi.NewInlineKeyboardButtonData("Cтать участником клуба на 2 месяц", "2Mount"),

				),
				tgbotapi.NewInlineKeyboardRow(

					tgbotapi.NewInlineKeyboardButtonData("Cтать участником клуба на 3 месяц", "3Mount"),
				),


			)
			msg.ReplyMarkup = numericKeyboard
			bot.Send(msg)
		case FREE:
			msg:= tgbotapi.NewMessage(update.Message.Chat.ID, "Вы можете полчить доступ к боту бесплатно, зарегистрировашись по нашей реферальной ссылке на binary.com и пополнив там торговый счет.")
			numericKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Получить ссылку для регистрации", GETLINK),
				),)
			msg.ReplyMarkup = numericKeyboard
			bot.Send(msg)


		default:
			if b, _:=regexp.MatchString(PATTERN, update.Message.Text ); b {
				UserID := update.Message.From.ID
				UserName := update.Message.From.UserName
				chatID := update.Message.Chat.ID
				str:=fmt.Sprint("/yes ", "`",UserID,"`", " 1-3(кол-во месяцев подписки)")
				err :=stor.AddUser(UserName,int(chatID), UserID)
				if err != nil {
					msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "Вы уже ожидаете активацию аккаунта")
					bot.Send(msg)

				}
				msg2:= tgbotapi.NewMessage(ADMINCHAT, update.Message.Text + " "+ str)
				msg2.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg2)
				msg:= tgbotapi.NewMessage(update.Message.Chat.ID, "Почтовый адрес отправлен на проверку")

				bot.Send(msg)
				continue
			}

			if update.Message.Photo != nil {
				Photo := *update.Message.Photo
				if len(Photo) != 0 {

					msg:= tgbotapi.NewPhotoShare(ADMINCHAT, Photo[0].FileID)
					bot.Send(msg)
					UserID := update.Message.From.ID
					UserName := update.Message.From.UserName
					chatID := update.Message.Chat.ID
					str:=fmt.Sprint("/yes ", "`",UserID,"`", " 1-3(кол-во месяцев подписки)")
					err :=stor.AddUser(UserName,int(chatID), UserID)
					if err != nil {
						msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "Вы уже ожидаете активацию аккаунта")
						bot.Send(msg)

					}
					msg2:= tgbotapi.NewMessage(ADMINCHAT, str)
					msg2.ParseMode = tgbotapi.ModeMarkdown
					bot.Send(msg2)
				}
				msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "После проверки оплаты вы получите доступ к сигналам")
				bot.Send(msg)
				continue
			}
			if update.Message.Chat.ID == ADMINCHAT {
				continue
			}
			msg:=tgbotapi.NewMessage(update.Message.Chat.ID, "Команда не валидна")
			bot.Send(msg)

		}


		fmt.Println("Test End")
	}

}
