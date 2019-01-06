package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"roman/client"
	"roman/database"
	rn "roman/proto/roman"
	grpcServer "roman/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var repository *client.Repository

func init() {
	database.AutoMigration()
	repository = client.NewRepository()
}

func runServer() {
	log.Println("Running on grpc server....")
	lis, err := net.Listen("tcp", ":3214")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC to port :3214")
	}
	s := grpc.NewServer()
	rn.RegisterRomanServer(s, grpcServer.NewGrpcServer())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// Running server GRPC with go routine
	go runServer()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputText := scanner.Text()
		repository.SetInputText(inputText)
		resultString := repository.Handle()
		if resultString != "" {
			fmt.Println("Result : " + resultString)
		}
	}
}
