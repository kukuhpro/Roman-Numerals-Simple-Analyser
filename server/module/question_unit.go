package module

import (
	"fmt"
	"roman/database"
	rn "roman/proto/roman"
	"strings"

	roman "github.com/StefanSchroeder/Golang-Roman"
)

type QuestionUnit struct {
	model          *database.Model
	listUnitNumber []database.UnitNumber
}

func (qu *QuestionUnit) generateListUnit(stringResult string, tokenAnalysis *rn.TokenAnalysis) string {
	if tokenAnalysis == nil {
		return stringResult
	}
	if tokenAnalysis.Key == KEY_UNIT {
		stringResult = stringResult + "," + tokenAnalysis.Value
	}
	return qu.generateListUnit(stringResult, tokenAnalysis.Next)
}

func (qu *QuestionUnit) SetListUnitNumber(listUnitName []string) {
	listUnitNumber := qu.model.GetListUnitNumber(listUnitName)
	qu.listUnitNumber = listUnitNumber
}

func (qu *QuestionUnit) GetRomanic(value string) string {
	for _, data := range qu.listUnitNumber {
		if data.Name == value {
			return data.RomanValue
		}
	}
	return ""
}

func (qu *QuestionUnit) GenerateRomanText(arrayUnitText []string) string {
	var romanText string
	for _, v := range arrayUnitText {
		romanText = romanText + qu.GetRomanic(v)
	}
	return romanText
}

func (qu *QuestionUnit) GetRomanicText(listUnitText string) string {
	arrayUnitText := strings.Split(listUnitText, ",")
	qu.SetListUnitNumber(arrayUnitText)
	return qu.GenerateRomanText(arrayUnitText)
}

func (qu *QuestionUnit) GetTotalAmount(tokenAnalysis *rn.TokenAnalysis) int {
	stringListUnit := qu.generateListUnit("", tokenAnalysis)
	amount := roman.Arabic(qu.GetRomanicText(stringListUnit))
	return amount
}

func (qu *QuestionUnit) GenerateResponseString(amount int) string {
	var result string
	for _, data := range qu.listUnitNumber {
		result = result + " " + data.Name
	}
	if result == "" {
		return result
	}
	return result + " is " + fmt.Sprint(amount)
}

func (qu *QuestionUnit) Process(tokenAnalysis *rn.TokenAnalysis) string {
	return qu.GenerateResponseString(qu.GetTotalAmount(tokenAnalysis))
}

func NewQuestionUnit() *QuestionUnit {
	var questionUnit QuestionUnit
	questionUnit.model = database.NewModel()
	return &questionUnit
}
