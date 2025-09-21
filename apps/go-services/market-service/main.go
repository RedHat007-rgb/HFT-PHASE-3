package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RedHat007-rgb/hft-phase-3/packages/golib/redis"
	"github.com/RedHat007-rgb/hft-phase-3/packages/golib/ws"
)


func main(){
	fmt.Println("Hello... from market-service");
	client:=redis.ConnectToRedis();
	ctx:=context.Background()
	if err:=client.Ping(ctx).Err();err!=nil{
		log.Fatal("error while connecting...")
	}
	// websocket....
	websocket:=ws.WsConnection();
	_,msg,err:=websocket.ReadMessage();
	if err!=nil{
		log.Println("unable to read messages....")
	}
	log.Println(msg)
}