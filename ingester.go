package main

import (
	"fmt"
	ws "github.com/gorilla/websocket"
	gdax "github.com/preichenberger/go-gdax"
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
		fmt.Println(message.Type)
		if message.Type == "ticker" {
			fmt.Println(message)
		}
	}
}

func NewIngester(conf *Configuration) *Ingester {
	i := new(Ingester)
	i.Conf = conf
	i.Url = "wss://ws-feed.gdax.com"

	// Postgres
	/*
		i.PostgresClient = NewPostgresClient(i.Conf.PGHost, i.Conf.PGPort,
			i.Conf.PGUser, i.Conf.PGPassword, i.Conf.PGDbname)
	*/

	return i
}
