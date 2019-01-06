package server

import (
	rn "roman/proto/roman"
)

type TokenAnalyser struct {
	tokenAnalysis *rn.TokenAnalysis
}

func (ta *TokenAnalyser) SetTokenAnalysis(tokenAnalysis *rn.TokenAnalysis) {
	ta.tokenAnalysis = tokenAnalysis
}

func (ta *TokenAnalyser) GetTokenAnalysis() *rn.TokenAnalysis {
	return ta.tokenAnalysis
}

func (ta *TokenAnalyser) generatePatternKey(tokenAnalysis *rn.TokenAnalysis) string {
	if tokenAnalysis == nil {
		return ""
	}
	tokenKey := "^" + tokenAnalysis.Key
	return tokenKey + ta.generatePatternKey(tokenAnalysis.Next)
}

func (ta *TokenAnalyser) GetTextKey() string {
	return ta.generatePatternKey(ta.tokenAnalysis)
}

func NewTokenAnalyser() *TokenAnalyser {
	var tokenAnalyser TokenAnalyser
	return &tokenAnalyser
}
