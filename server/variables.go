package server

import (
	rn "roman/proto/roman"
	"roman/server/module"
)

type RequestProcessValidTextKey struct {
	TextKey       string
	TokenAnalysis *rn.TokenAnalysis
	DefineModule  string
}

type ModuleContract interface {
	Process(tokenAnalysis *rn.TokenAnalysis) string
}

var (
	DEFINE_MATERIAL          = "define_material"
	DEFINE_UNIT              = "define_unit"
	DEFINE_QUESTION_UNIT     = "define_question_unit"
	DEFINE_QUESTION_MATERIAL = "define_question_material"

	PATTERN_DEFINE_MATERIAL   = `(\^unit)+\^material\^verb\^numeric\^credit`
	PATTERN_DEFINE_UNIT       = `\^unit\^verb\^roman`
	PATTERN_QUESTION_UNIT     = `\^question\^quantifiers\^verb(\^unit)+\^mark`
	PATTERN_QUESTION_MATERIAL = `\^question\^quantifiers\^credit\^verb(\^unit)+\^material\^mark`

	listPatternKeyText = map[string]string{
		DEFINE_MATERIAL:          PATTERN_DEFINE_MATERIAL,
		DEFINE_UNIT:              PATTERN_DEFINE_UNIT,
		DEFINE_QUESTION_UNIT:     PATTERN_QUESTION_UNIT,
		DEFINE_QUESTION_MATERIAL: PATTERN_QUESTION_MATERIAL,
	}

	listModuleProcess = map[string]ModuleContract{
		DEFINE_MATERIAL:          module.NewMaterial(),
		DEFINE_QUESTION_MATERIAL: module.NewQuestionMaterial(),
		DEFINE_QUESTION_UNIT:     module.NewQuestionUnit(),
		DEFINE_UNIT:              module.NewUnit(),
	}
)
