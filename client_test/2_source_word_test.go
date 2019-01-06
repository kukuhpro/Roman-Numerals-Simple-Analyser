package client_test

import (
	"roman/client"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SourceWordSuiteTest struct {
	sourceWord *client.SourceWord
}

var sourceWordSuiteTest SourceWordSuiteTest

func init() {
	sourceWordSuiteTest.sourceWord = client.NewSourceWord()
}

func TestWordSetterGetter(t *testing.T) {
	sourceWord := sourceWordSuiteTest.sourceWord
	word := "Hello"
	sourceWord.SetWord(word)

	resultWord := sourceWord.GetWord()
	assert.Equal(t, word, resultWord)
}

func TestIsMatchPattern(t *testing.T) {
	sourceWord := sourceWordSuiteTest.sourceWord
	pattern := "[0-9]*"
	numericString := "2873620"
	sourceWord.SetWord(numericString)
	expected := true
	result := sourceWord.IsMatchPattern(pattern)
	assert.Equal(t, expected, result)
}

func TestGetSource(t *testing.T) {
	sourceWord := sourceWordSuiteTest.sourceWord
	expected := client.KEY_MATERIAL
	word := "Gold"
	sourceWord.SetWord(word)
	result := sourceWord.GetSource()
	assert.Equal(t, expected, result)
}
