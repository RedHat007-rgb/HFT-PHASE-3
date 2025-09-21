package ws

import (
	"log"

	"github.com/gorilla/websocket"
)



func WsConnection() *websocket.Conn{
	conn,_,err:=websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/btcusdt@ticker",nil);
	if err!=nil{
		log.Fatal("error connecting ws...",err)
	}

	return conn
}