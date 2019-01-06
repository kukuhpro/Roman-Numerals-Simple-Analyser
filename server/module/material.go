package module

import (
	"roman/database"
	rn "roman/proto/roman"
	"strconv"
)

type Material struct {
	model        *database.Model
	questionUnit *QuestionUnit
}

func (m *Material) GetMaterialText(tokenAnalysis *rn.TokenAnalysis) string {
	if tokenAnalysis == nil {
		return ""
	} else if tokenAnalysis.Key == KEY_MATERIAL {
		return tokenAnalysis.Value
	}
	return m.GetMaterialText(tokenAnalysis.Next)
}

func (m *Material) GetNumericText(tokenAnalysis *rn.TokenAnalysis) string {
	if tokenAnalysis == nil {
		return ""
	} else if tokenAnalysis.Key == KEY_NUMERIC {
		return tokenAnalysis.Value
	}
	return m.GetNumericText(tokenAnalysis.Next)
}

func (m *Material) GetRealValueCredit(tokenAnalysis *rn.TokenAnalysis) float64 {
	amount := m.questionUnit.GetTotalAmount(tokenAnalysis)
	creditValue := m.GetNumericText(tokenAnalysis)
	i, _ := strconv.ParseFloat(creditValue, 64)
	return i / float64(amount)
}

func (m *Material) GenerateStoreMaterial(materialName string, materialValueCredit float64) database.MaterialCredit {
	var dataMaterial database.MaterialCredit
	dataMaterial.Name = materialName
	dataMaterial.CreditValue = materialValueCredit
	return dataMaterial
}

func (m *Material) CreateOrUpdateMaterial(materialName string, materialValueCredit float64) {
	detailMaterial := m.model.DetailMaterial(materialName)
	storeMaterialData := m.GenerateStoreMaterial(materialName, materialValueCredit)
	if detailMaterial.ID == 0 {
		m.model.CreateMaterial(storeMaterialData)
	} else {
		m.model.UpdateMaterial(detailMaterial.ID, storeMaterialData)
	}
}

func (m *Material) Process(tokenAnalysis *rn.TokenAnalysis) string {
	materialName := m.GetMaterialText(tokenAnalysis)
	materialValueCredit := m.GetRealValueCredit(tokenAnalysis)
	m.CreateOrUpdateMaterial(materialName, materialValueCredit)
	return ""
}

func NewMaterial() *Material {
	var material Material
	material.questionUnit = NewQuestionUnit()
	material.model = database.NewModel()
	return &material
}
