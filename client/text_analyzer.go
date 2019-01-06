package client

import (
	rn "roman/proto/roman"
	"strings"
)

type TextAnalyzer struct {
	inputText  string
	sourceWord *SourceWord
}

func (ta *TextAnalyzer) SetInputText(inputText string) {
	ta.inputText = inputText
}

func (ta *TextAnalyzer) GetInputText() string {
	return ta.inputText
}

func (ta *TextAnalyzer) SplitText() []string {
	return strings.Split(ta.inputText, " ")
}

func (ta *TextAnalyzer) GetKeyToken(word string) string {
	sourceWord := ta.sourceWord
	sourceWord.SetWord(word)
	return sourceWord.GetSource()
}

func (ta *TextAnalyzer) CreateTokenAnalysis(word string) *rn.TokenAnalysis {
	var tokenAnalysis rn.TokenAnalysis
	tokenAnalysis.Key = ta.GetKeyToken(word)
	tokenAnalysis.Value = word
	return &tokenAnalysis
}

func (ta *TextAnalyzer) GenerateTokenAnalysis(index int) *rn.TokenAnalysis {
	arrayString := ta.SplitText()
	if index >= len(arrayString) {
		return nil
	}
	tokenAnalysis := ta.CreateTokenAnalysis(arrayString[index])
	nextTokenAnalysis := ta.GenerateTokenAnalysis(index + 1)
	if nextTokenAnalysis != nil {
		tokenAnalysis.Next = nextTokenAnalysis
	}
	return tokenAnalysis
}

func (ta *TextAnalyzer) Analysis() *rn.TokenAnalysis {
	tokenAnalysis := ta.GenerateTokenAnalysis(0)
	return tokenAnalysis
}

func NewTextAnalyzer() *TextAnalyzer {
	var analyser TextAnalyzer
	analyser.sourceWord = NewSourceWord()
	return &analyser
}
