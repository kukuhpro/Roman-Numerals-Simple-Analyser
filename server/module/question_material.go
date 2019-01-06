package module

import (
	"fmt"
	"roman/database"
	rn "roman/proto/roman"
)

type QuestionMaterial struct {
	model        *database.Model
	questionUnit *QuestionUnit
	material     *Material
}

func (qm *QuestionMaterial) GetDetailMaterial(materialName string) database.MaterialCredit {
	return qm.model.DetailMaterial(materialName)
}

func (qm *QuestionMaterial) GetCreditValue(materialName string) float64 {
	detailMaterial := qm.GetDetailMaterial(materialName)
	return detailMaterial.CreditValue
}

func (qm *QuestionMaterial) GetStringUnitMaterial(tokenAnalysis *rn.TokenAnalysis) string {
	if tokenAnalysis == nil {
		return ""
	} else if tokenAnalysis.Key == KEY_UNIT || tokenAnalysis.Key == KEY_MATERIAL {
		return tokenAnalysis.Value + " " + qm.GetStringUnitMaterial(tokenAnalysis.Next)
	}
	return qm.GetStringUnitMaterial(tokenAnalysis.Next)
}

func (qm *QuestionMaterial) GenerateResponseString(tokenAnalysis *rn.TokenAnalysis) string {
	return qm.GetStringUnitMaterial(tokenAnalysis)
}

func (qm *QuestionMaterial) GetTotalCreditMaterial(tokenAnalysis *rn.TokenAnalysis) float64 {
	creditValueMaterial := qm.GetCreditValue(qm.material.GetMaterialText(tokenAnalysis))
	amountUnit := qm.questionUnit.GetTotalAmount(tokenAnalysis)

	return creditValueMaterial * float64(amountUnit)
}
func (qm *QuestionMaterial) Process(tokenAnalysis *rn.TokenAnalysis) string {
	totalCreditMaterial := qm.GetTotalCreditMaterial(tokenAnalysis)
	responseStringUnitMaterial := qm.GenerateResponseString(tokenAnalysis)
	return responseStringUnitMaterial + " is " + fmt.Sprint(totalCreditMaterial)
}

func NewQuestionMaterial() *QuestionMaterial {
	var questionMaterial QuestionMaterial
	questionMaterial.questionUnit = NewQuestionUnit()
	questionMaterial.model = database.NewModel()
	return &questionMaterial
}
