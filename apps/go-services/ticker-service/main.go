package main

import (
	"log"
	"net"

	tickerpb "github.com/RedHat007-rgb/hft-phase-3/packages/golib/proto/ticker"
	"google.golang.org/grpc"
)
type server struct {
tickerpb.UnimplementedTickerServiceServer
}
func (s *server) StreamTicker(req *tickerpb.TickerRequest, stream tickerpb.TickerService_StreamTickerServer) error {
// TODO: Implement your streaming logic here.
// For example, loop to send updates:
for {
update := &tickerpb.TickerUpdate{
Exchange: "example",
Symbol:   req.Symbol,
Price:    "100.00",
Volume:   "1000",
EventTime: 1234567890,
}
if err := stream.Send(update); err != nil {
return err
}
// Add delay or real data fetching logic.
}
return nil
}
func main() {
lis, err := net.Listen("tcp", ":50051") // Change port as needed.
if err != nil {
log.Fatalf("failed to listen: %v", err)
}
s := grpc.NewServer()
tickerpb.RegisterTickerServiceServer(s, &server{})
log.Printf("gRPC server listening on :50051")
if err := s.Serve(lis); err != nil {
log.Fatalf("failed to serve: %v", err)
}
}
