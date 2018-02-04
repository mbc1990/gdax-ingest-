package main

import (
	"fmt"
	ws "github.com/gorilla/websocket"
	gdax "github.com/preichenberger/go-gdax"
	"strconv"
)

type Ingester struct {
	PostgresClient *PostgresClient
	Conf           *Configuration
	Url            string
}

func (i *Ingester) Start() {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial(i.Url, nil)
	if err != nil {
		println(err.Error())
	}

	subscribe := gdax.Message{
		Type: "subscribe",
		Channels: []gdax.MessageChannel{
			gdax.MessageChannel{
				Name: "ticker",
				ProductIds: []string{
					"ETH-USD",
				},
			},
		},
	}
	if err := wsConn.WriteJSON(subscribe); err != nil {
		println(err.Error())
	}

	message := gdax.Message{}
	for true {
		if err := wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}
		if message.Type == "ticker" {
			txnTime := message.Time.Time()

			// For some reason, the first message we get has negative time and is
			// missing information, so just throw that away for now
			if txnTime.Unix() < 0 {
				continue
			}
			ticksCounter.Inc()
			price := message.Price
			side := message.Side
			if side == "buy" {
				buyGauge.Set(price)
			} else {
				sellGauge.Set(price)
			}
			fmt.Println("Time: " + strconv.Itoa(int(txnTime.Unix())))
			fmt.Println("Price: ", price)
			fmt.Println("Side: " + side)
			fmt.Println("-------------------------")

			i.PostgresClient.InsertTick(side, price, int(txnTime.Unix()))
		}
	}
}

func NewIngester(conf *Configuration) *Ingester {
	i := new(Ingester)
	i.Conf = conf
	i.Url = "wss://ws-feed.gdax.com"

	// Postgres
	i.PostgresClient = NewPostgresClient(i.Conf.PGHost, i.Conf.PGPort,
		i.Conf.PGUser, i.Conf.PGPassword, i.Conf.PGDbname)

	return i
}
