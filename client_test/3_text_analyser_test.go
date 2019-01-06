package client_test

import (
	"fmt"
	"roman/client"
	"testing"

	rn "roman/proto/roman"

	"github.com/stretchr/testify/assert"
)

type TextAnalyzerSuiteTest struct {
	textAnalyser *client.TextAnalyzer
}

var textAnalyzerSuiteTest TextAnalyzerSuiteTest

func init() {
	textAnalyzerSuiteTest.textAnalyser = client.NewTextAnalyzer()
}

var tableDrivenTestSet = map[string]string{
	"glip is 45": "glip is 45",
}

var tableDrivenRomanAnalyzer = map[string]*rn.TokenAnalysis{
	"glob is I": &rn.TokenAnalysis{
		Value: "glob",
		Key:   "unit",
		Next:  &rn.TokenAnalysis{},
	},
}

func TestSetterGetterText(t *testing.T) {
	textAnalyser := textAnalyzerSuiteTest.textAnalyser
	for k, _ := range tableDrivenTestSet {
		inputText := k
		textAnalyser.SetInputText(inputText)
		assert.Equal(t, inputText, textAnalyser.GetInputText())
	}

}

func TestSplitText(t *testing.T) {
	textAnalyser := textAnalyzerSuiteTest.textAnalyser
	expectedResult := []string{"hello", "world"}
	inputText := "hello world"
	textAnalyser.SetInputText(inputText)
	result := textAnalyser.SplitText()

	assert.IsType(t, []string{}, result)
	assert.EqualValues(t, expectedResult, result)
}

func TestKeyToken(t *testing.T) {
	textAnalyser := textAnalyzerSuiteTest.textAnalyser
	expectedResult := client.KEY_MATERIAL
	word := "silver"
	result := textAnalyser.GetKeyToken(word)
	assert.Equal(t, expectedResult, result)
}

func TestCreateTokenAnalysis(t *testing.T) {
	textAnalyser := textAnalyzerSuiteTest.textAnalyser
	word := "silver"
	expectedKey := client.KEY_MATERIAL
	tokenAnalysis := textAnalyser.CreateTokenAnalysis(word)
	assert.Equal(t, expectedKey, tokenAnalysis.Key)
	assert.Equal(t, word, tokenAnalysis.Value)
}

func TestGenerateTokenAnalysis(t *testing.T) {
	textAnalyser := textAnalyzerSuiteTest.textAnalyser
	inputText := "glob is I"
	textAnalyser.SetInputText(inputText)
	tokenAnalysis := textAnalyser.Analysis()
	nextFourth := tokenAnalysis.Next.Next.Next
	nextThird := tokenAnalysis.Next.Next
	assert.Equal(t, "<nil>", fmt.Sprint(nextFourth))
	assert.Equal(t, client.KEY_ROMAN, nextThird.Key)
}
