package main

import (
	"encoding/json"
	"log"
	"net"

	pb "github.com/RedHat007-rgb/hft-phase-3/apps/go-services/ticker-service/proto/ticker"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type TickerServer struct{
	pb.UnimplementedTickerServiceServer
}


type BinanceTicker struct{
	
	Symbol string `json:"s"`
	
	Volume string `json:"v"`
	High string `json:"h"`
}

func (s *TickerServer) StreamTicker(req *pb.TickerRequest,stream pb.TickerService_StreamTickerServer) error{
	url:="wss://stream.binance.com:9443/ws/"+req.Symbol+"@ticker"
	conn,_,err:=websocket.DefaultDialer.Dial(url,nil)

	if err!=nil{
		return err
	}

	defer conn.Close()


	log.Println("subscribed to ",req.Symbol)

	for{
		_,msg,err:=conn.ReadMessage()
		if err!=nil{
			log.Println("read error",err)
		}

		var t BinanceTicker


		if err:=json.Unmarshal(msg,&t);err!=nil{
			// log.Println(json.Unmarshal(msg,&t))
			log.Println("unmarshall error",err)
		}
		
		update:=&pb.TickerUpdate{
			Symbol: t.Symbol,
			
			Volume: t.Volume,
			
			High:t.High,
			
		}

		log.Println(update)
		if err:=stream.Send(update);err!=nil{
			return err
		}
	}

}


func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("❌ failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTickerServiceServer(grpcServer, &TickerServer{})
	log.Println("🚀 Go Ticker Service running on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ failed to serve: %v", err)
	}
}