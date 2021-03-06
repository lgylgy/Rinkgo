package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/lgylgy/rinkgo/pkg/services/pstat"
	pb "github.com/lgylgy/rinkgo/pkg/services/pstat/proto"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 5000, "The server port")
	url := flag.String("url", "", "The database url")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPStatServiceServer(grpcServer, pstat.NewPStatServer(*url))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
