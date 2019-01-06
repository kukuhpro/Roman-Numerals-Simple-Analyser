package module

import (
	"roman/database"
	rn "roman/proto/roman"
)

type Unit struct {
	model         *database.Model
	tokenAnalysis *rn.TokenAnalysis
}

func (u *Unit) GetUnitText(tokenAnalysis *rn.TokenAnalysis) string {
	if tokenAnalysis.Key == KEY_UNIT {
		return tokenAnalysis.Value
	}
	return u.GetUnitText(tokenAnalysis.Next)
}

func (u *Unit) GetRomanText(tokenAnalysis *rn.TokenAnalysis) string {
	if tokenAnalysis.Key == KEY_ROMAN {
		return tokenAnalysis.Value
	}
	return u.GetRomanText(tokenAnalysis.Next)
}

func (u *Unit) GenerateStoreUnit(unitText, romanText string) database.UnitNumber {
	var storeUnit database.UnitNumber
	storeUnit.Name = unitText
	storeUnit.RomanValue = romanText
	return storeUnit
}

func (u *Unit) CreateOrUpdateUnit(unitText, romanText string) database.UnitNumber {
	detailUnit := u.model.DetailUnit(unitText)
	storeUnit := u.GenerateStoreUnit(unitText, romanText)
	if detailUnit.ID != 0 {
		u.model.UpdateUnit(detailUnit.ID, storeUnit)
	} else {
		u.model.CreateUnit(storeUnit)
	}
	return storeUnit
}

func (u *Unit) Process(tokenAnalysis *rn.TokenAnalysis) string {
	unitText := u.GetUnitText(tokenAnalysis)
	romanText := u.GetRomanText(tokenAnalysis)
	u.CreateOrUpdateUnit(unitText, romanText)
	return ""
}

func NewUnit() *Unit {
	var unit Unit
	unit.model = database.NewModel()
	return &unit
}
