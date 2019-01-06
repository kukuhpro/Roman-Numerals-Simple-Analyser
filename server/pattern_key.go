package server

import (
	"regexp"
)

type PatternKey struct {
	textKey string
}

func (pk *PatternKey) SetTextKey(textKey string) {
	pk.textKey = textKey
}

func (pk *PatternKey) GetTextKey() string {
	return pk.textKey
}

func (pk *PatternKey) IsMatchPattern(pattern string) bool {
	r, _ := regexp.Compile(pattern)
	return r.MatchString(pk.textKey)
}

func (pk *PatternKey) LoopListPatternKey() string {
	for k, v := range listPatternKeyText {
		if pk.IsMatchPattern(v) {
			return k
		}
	}
	return ""
}

func (pk *PatternKey) GetPatternDefine() string {
	return pk.LoopListPatternKey()
}

func NewPatternKey() *PatternKey {
	var patternKey PatternKey
	return &patternKey
}
