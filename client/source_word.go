package client

import (
	"regexp"
)

type SourceWord struct {
	word string
}

func (st *SourceWord) SetWord(word string) {
	st.word = word
}

func (st *SourceWord) GetWord() string {
	return st.word
}

func (st *SourceWord) IsMatchPattern(pattern string) bool {
	r, _ := regexp.Compile(pattern)
	return r.MatchString(st.word)
}

func (st *SourceWord) GetSource() string {
	var source string
	for k, v := range regexPatternSourceWord {
		if st.IsMatchPattern(v) {
			source = k
			break
		}
	}
	return source
}

func NewSourceWord() *SourceWord {
	var sourceWord SourceWord
	return &sourceWord
}
