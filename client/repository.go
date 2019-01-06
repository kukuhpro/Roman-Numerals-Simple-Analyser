package client

import (
	"context"
	rn "roman/proto/roman"
	"time"
)

type Repository struct {
	inputText    string
	textAnalyzer *TextAnalyzer
	client       *GrpcClient
}

func (r *Repository) SetInputText(inputText string) {
	r.inputText = inputText
}

func (r *Repository) GetInputText() string {
	return r.inputText
}

func (r *Repository) GetTokenAnalysis() *rn.TokenAnalysis {
	textAnalyzer := r.textAnalyzer
	textAnalyzer.SetInputText(r.inputText)
	return textAnalyzer.Analysis()
}

func (r *Repository) ProcessAnalysis(tokenAnalysis *rn.TokenAnalysis) *rn.Response {
	var response *rn.Response
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.client.ctx = ctx
	response = r.client.ProcessAnalysis(tokenAnalysis)
	return response
}

func (r *Repository) MappingResponseAnalysis(response *rn.Response) string {
	var result string
	if response.Status == 1 {
		result = response.Response.Result
	} else {
		result = response.Error.ErrorText
	}
	return result
}

func (r *Repository) Handle() string {
	tokenAnalysis := r.GetTokenAnalysis()
	return r.MappingResponseAnalysis(r.ProcessAnalysis(tokenAnalysis))
}

func NewRepository() *Repository {
	var repository Repository
	repository.textAnalyzer = NewTextAnalyzer()
	repository.client = NewGrpcClient()
	return &repository
}
