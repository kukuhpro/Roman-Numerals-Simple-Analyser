package server

import (
	"context"
	rs "roman/proto/roman"
)

type GrpcServer struct {
	repository *Repository
}

func (gs *GrpcServer) ProcessAnalysis(ctx context.Context, token *rs.TokenAnalysis) (*rs.Response, error) {
	var response *rs.Response
	response = gs.repository.ProcessAnalysis(token)
	return response, nil
}

func NewGrpcServer() *GrpcServer {
	var server GrpcServer
	server.repository = NewRepository()
	return &server
}
