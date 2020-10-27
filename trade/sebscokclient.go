package trade

import (
	"TelegrammBOTOPTIONS/botApi"
	"TelegrammBOTOPTIONS/store"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
)

var GMT, _ = time.LoadLocation("Africa/Abidjan")
const TIMETICK = 60


type SignalsPool struct {
	Signals map[int64]Signal
	StartSignals map[int64]Signal
	m sync.Mutex
}
//REQUST GROUP
type reqid struct {
	Req int `json:"req_id"`
}
type Passthrough struct {
	Passthrough string `json:"passthrough"`
}
type ActiveSymbols struct {
	ActiveSymbols string `json:"active_symbols"` //brief
}
type req1 struct {
	Subscribe int `json:"subscribe"`
	Passthrough
	Website_status int `json:"website_status"`
	reqid
}
type req2 struct {
	Time int `json:"time"`
	Passthrough
	reqid
}
type req3 struct {
	reqid
	Landing_company string `json:"landing_company"` //ru
}
type payoutcurrencies struct {
	PayoutCurrencies int `json:"payout_currencies"` // 1
}

type req4 struct {
	Passthrough
	forgetall
	reqid

}
type req5 struct {
	Forgetall string `json:"forget_all"`
	reqid
	Passthrough
}
type req6 struct {
	Forgetall string `json:"forget_all"`
	reqid
	Passthrough
}
type contractsfor struct {
	ContractsFor string `json:"contracts_for"`
}
type forgetall struct {
	Forgetall []string `json:"forget_all"`
}
type Resp struct {
	Tick Tick `json:"tick"`
}
type Tick struct {
	Ask float32 `json:"ask"` // 8462.93
	Bid float32 `json:"bid"` // 8460.93
	Epoch int64 `json:"epoch"` //1601199397
	Quote float32 `json:"quote"`// 8461.93

}
type Signal struct {
	TimeStart int64
	TimeEnd int64
	Price float32
	Rise bool
	Text string
}
func NewSignal(timeStart int64, rise bool, Price float32) Signal {
	s:= Signal{}

	s.Rise = rise
	s.TimeStart = timeStart+TIMETICK
	s.TimeEnd = timeStart+TIMETICK+TIMETICK
	s.Price = Price
	if s.Rise {
		s.Text = "ВВЕРХ"
	} else {
		s.Text = "ВНИЗ"
	}
	return s
}
func NewSignalsPool() SignalsPool {
	s:= SignalsPool{}
	s.Signals =make(map[int64]Signal, 10)
	s.StartSignals =make(map[int64]Signal, 10)
	return s
}
func(s *SignalsPool) AddNewSignal(si Signal) bool {
	if _, ok:=s.StartSignals[si.TimeStart]; !ok {
		s.m.Lock()
		s.StartSignals[si.TimeStart] = si
		s.m.Unlock()
		return true
	}
	return false
}
//true если сигнал отработал
func(s *SignalsPool) CheckSignalEnd(TimeEnd int64, Quote float32) (bool, Signal) {

	if sig, ok := s.Signals[TimeEnd]; ok {
		s.m.Lock()
		delete(s.Signals, TimeEnd)
		s.m.Unlock()
		switch {
		case Quote>sig.Price && sig.Rise:
			return  true, sig
		case Quote<sig.Price && !sig.Rise:
			return true, sig
		default:
			return false, sig

		}

	}
	return false, Signal{}
}
func(s *SignalsPool) CheckSignalStart(TimeStart int64, Quote float32, p *botApi.PoolChats) {

	if sig, ok := s.StartSignals[TimeStart]; ok {
		s.m.Lock()
		delete(s.StartSignals, TimeStart)
		sig.Price = Quote
		s.Signals[sig.TimeEnd] = sig
		s.m.Unlock()
		 f:=strconv.FormatFloat(float64(Quote), 'G', -1, 64 )
		text :="Старт сигнала, цена: "+ f
		p.SendMessage(text)
	}

}
func (s *Signal) SendResult(result bool, Quote float32, p *botApi.PoolChats) {
	f:=strconv.FormatFloat(float64(Quote), 'G', -1, 64 )
	text := ""
	if result {
		text = "Сигнал отработал " +f
	} else {
		text = "Сигнал не отработал " +f
	}
	p.SendMessage(text)

}

const swconn =  "wss://blue.binaryws.com/websockets/v3?app_id=1&l=EN"
func ConnectBinary(signal, stopsignal chan int, bot *tgbotapi.BotAPI, stor *store.MySQL) {
	c, _, err := websocket.DefaultDialer.Dial(swconn, nil)
	if err != nil {
		fmt.Println(err)
	}
	 func() {
		 r1:=[]byte(`{"authorize":"a1-aIfqSYsjkdq1NNbg2DzSfwjLLMoNk","req_id":1,"passthrough":{}}`)
		 r2:=[]byte(`{"website_status":1,"subscribe":1,"req_id":2,"passthrough":{}}`)
		 r3:=[]byte(`{"time":1,"req_id":3,"passthrough":{}}`)
		 r4:=[]byte(`{"balance":1,"subscribe":1,"req_id":4,"passthrough":{}}`)
		 r5:=[]byte(`{"get_settings":1,"req_id":5,"passthrough":{}}`)
		 r6:=[]byte(`{"get_account_status":1,"req_id":6,"passthrough":{}}`)
		 r7:=[]byte(`{"payout_currencies":1}`)
		 r8:=[]byte(`{"mt5_login_list":1,"req_id":7,"passthrough":{}}`)
		 r9:=[]byte(`{"transaction":1,"subscribe":1,"req_id":8,"passthrough":{}}`)
		 r10:=[]byte(`{"landing_company":"ru","req_id":9,"passthrough":{}}`)
		 r11:=[]byte(`{"payout_currencies":1}`)
		 r12:= []byte(`{"active_symbols":"brief"}`)
		 r13:= []byte(`{"forget_all":["ticks","candles"],"req_id":10,"passthrough":{}}`)
		 r14:= []byte(`{"contracts_for":"frxEURUSD"}`)
		 r15:= []byte(`{"forget_all":"proposal_open_contract","req_id":11,"passthrough":{}}`)
		 r16:= []byte(`{"forget_all":"proposal","req_id":12,"passthrough":{}}`)
		 r17:= []byte(`{"statement":1,"limit":1,"req_id":13,"passthrough":{}}`)
		 r18:= []byte(`{"ticks_history":"frxEURUSD","style":"ticks","end":"latest","count":20,"subscribe":1,"req_id":14,"passthrough":{}}`)
		 r19:= []byte(`{"proposal":1,"subscribe":1,"amount":10,"basis":"stake","contract_type":"CALL","currency":"USD","symbol":"frxEURUSD","duration":5,"duration_unit":"t","passthrough":{"form_id":1},"req_id":15}`)
	     r20:= []byte(`{"proposal":1,"subscribe":1,"amount":10,"basis":"stake","contract_type":"PUT","currency":"USD","symbol":"frxEURUSD","duration":5,"duration_unit":"t","passthrough":{"form_id":1},"req_id":16}`)

		 //rr1, _:=json.Marshal(r1)
		//rr2, _:=json.Marshal(r2)
		//rr3, _:=json.Marshal(r3)
		//rr4, _:=json.Marshal(r4)
		c.WriteMessage(1,r1)
		c.WriteMessage(1,r2)
		c.WriteMessage(1,r3)
		 c.WriteMessage(1,r4)
		 c.WriteMessage(1,r5)
		 c.WriteMessage(1,r6)
		 c.WriteMessage(1,r7)
		 c.WriteMessage(1,r8)
		 c.WriteMessage(1,r9)
		 c.WriteMessage(1,r10)
		 c.WriteMessage(1,r11)
		 c.WriteMessage(1,r12)
		 c.WriteMessage(1,r13)
		 c.WriteMessage(1,r14)
		 c.WriteMessage(1,r15)
		 c.WriteMessage(1,r16)
		 c.WriteMessage(1,r17)
		 c.WriteMessage(1,r18)
		 c.WriteMessage(1,r19)
		 c.WriteMessage(1,r20)

         SignalAnalitic := NewQueue()
         PoolChat := botApi.NewPool(bot, 60)
         var sig Signal
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				go ConnectBinary(signal, stopsignal, bot, stor)
				PoolChat.SendMessage("Рестарт сервиса, для подключения к сигналам введите "+ botApi.GETSIG)
				return
			}

			resp := Resp{}
			json.Unmarshal(message, &resp)
			if resp.Tick.Epoch == 0 {
				continue
			}
			SignalAnalitic.Add(float64(resp.Tick.Quote))
			if sig.TimeStart != 0 || sig.TimeEnd != 0 {
				if sig.TimeStart <= resp.Tick.Epoch {
					if sig.TimeStart != 0 {
						sig.Price = resp.Tick.Quote
						f:=strconv.FormatFloat(float64(resp.Tick.Quote), 'G', -1, 64 )
						text :="Старт сигнала, цена:"+f[:7]
						PoolChat.SendMessage(text)
						sig.TimeStart = 0
					}

				}
				if sig.TimeEnd <= resp.Tick.Epoch {
					if  sig.TimeEnd != 0{
						f:=strconv.FormatFloat(float64(resp.Tick.Quote), 'G', -1, 64 )
						Quote := resp.Tick.Quote
						text := ""
						switch {
						case Quote>sig.Price && sig.Rise:
							text = "Отработал"
						case Quote<sig.Price && !sig.Rise:
							text = "Отработал"
						default:
							text = "Не отработал"

						}
						msg := "Сигнал " + text + ". Цена:" +f[:7]
						PoolChat.SendMessage(msg)
						sig = Signal{}
					}

				}
			}


			fmt.Println(resp.Tick.Quote, "||", resp.Tick.Epoch, PoolChat, sig)
			select {
			case idchat := <- signal:
				msg:= tgbotapi.NewMessage(int64(idchat), "Вы подписались на сигналы, таймфрейм 1М")
				bot.Send(msg)
				PoolChat.AddChat(idchat)
			case  <-PoolChat.Signal:
				fmt.Println("TIMER")
				t:=time.Now().In(GMT)
				t = t.Round(1*time.Minute)
				if sig.TimeEnd == 0 {
					var result bool
					// Вставляем сигнал
					result = SignalAnalitic.GetSolving()

					sig = NewSignal(t.Unix(), result,resp.Tick.Quote)

					t2:= time.Unix(sig.TimeStart, 0).In(GMT)
					fmt.Println(t2, sig)
					hour, minute, _:= t2.Clock()
					h:=strconv.Itoa(hour)
					if len(h)<2{
						h = "0"+h
					}
					m:=strconv.Itoa(minute)
					if len(m)<2 {
						m = "0"+m
					}
					text := "EUR/USD/" + sig.Text + "/" +h+":"+m+"GMT"
					PoolChat.SendMessage(text)
				}
			case idchat := <- stopsignal:
				fmt.Println(idchat, "STOPSIGNAL")
				PoolChat.OffChat(idchat)
				msg:= tgbotapi.NewMessage(int64(idchat), "Вы отключились от сигналов, для подключения введите " + botApi.GETSIG)
				bot.Send(msg)


			default:

				continue
			}




		}


	}()

}


//{"echo_req":{"passthrough":{},"req_id":3,"time":1},"msg_type":"time","passthrough":{},"req_id":3,"time":1601916930}

type ResponserECHO struct {
	Echo_req Echo_Req `json:"echo_req"`
	Msg_type string `json:"msg_type"`
	Passthrough
	reqid
	Time int `json:"time"`
}
type Echo_Req struct {
	Passthrough
	reqid
	Time int `json:"time"`
}

// {"authorize":{"account_list":[{"currency":"USD","is_disabled":0,"is_virtual":1,"landing_company_name":"virtual","loginid":"VRTC3442909"}],"balance":9995,"country":"ru","currency":"USD","email":"golem28@gmail.com","fullname":"  ","is_virtual":1,"landing_company_fullname":"Deriv Limited","landing_company_name":"virtual","local_currencies":{"RUB":{"fractional_digits":2}},"loginid":"VRTC3442909","scopes":["read","admin","trade","payments"],"upgradeable_landing_companies":["svg"],"user_id":7555752},"echo_req":{"authorize":"<not shown>","passthrough":{},"req_id":1},"msg_type":"authorize","passthrough":{},"req_id":1}