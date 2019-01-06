package client

import (
	"context"
	"log"
	rn "roman/proto/roman"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	address string
	ctx     context.Context
}

func (gc *GrpcClient) ProcessAnalysis(tokenAnalysis *rn.TokenAnalysis) *rn.Response {
	var response *rn.Response
	conn, err := grpc.Dial(gc.address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return response
	}
	defer conn.Close()
	c := rn.NewRomanClient(conn)
	response, err = c.ProcessAnalysis(gc.ctx, tokenAnalysis)
	if err != nil {
		log.Println(err)
		return response
	}
	return response
}

func NewGrpcClient() *GrpcClient {
	var client GrpcClient
	client.address = "localhost:3214"
	return &client
}
