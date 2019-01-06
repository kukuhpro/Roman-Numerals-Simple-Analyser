package server

import (
	rn "roman/proto/roman"
)

type Repository struct {
	tokenAnalyser *TokenAnalyser
	patternKey    *PatternKey
}

func (r *Repository) GetTextKey(tokenAnalysis *rn.TokenAnalysis) string {
	r.tokenAnalyser.SetTokenAnalysis(tokenAnalysis)
	return r.tokenAnalyser.GetTextKey()
}

func (r *Repository) InvalidFormatText() *rn.Response {
	var errorResponse rn.ErrorResponse
	errorResponse.ErrorText = "I Have no idea what are you talking about"
	var response rn.Response
	response.Status = 99
	response.Error = &errorResponse
	return &response
}

func (r *Repository) ValidFormatText(requestProcessValidTextKey RequestProcessValidTextKey) *rn.Response {
	var moduleContract ModuleContract
	moduleContract = listModuleProcess[requestProcessValidTextKey.DefineModule]
	var successResponse rn.SuccessResponse
	successResponse.Result = moduleContract.Process(requestProcessValidTextKey.TokenAnalysis)

	var response rn.Response
	response.Status = 1
	response.Response = &successResponse
	return &response
}

func (r *Repository) DefineModulePattern(textKey string) (string, bool) {
	r.patternKey.SetTextKey(textKey)
	patternDefine := r.patternKey.GetPatternDefine()
	return patternDefine, patternDefine != ""
}

func (r *Repository) ProcessAnalysis(tokenAnalysis *rn.TokenAnalysis) *rn.Response {
	textKey := r.GetTextKey(tokenAnalysis)
	patternDefine, flag := r.DefineModulePattern(textKey)
	if !flag {
		return r.InvalidFormatText()
	}
	var request RequestProcessValidTextKey
	request.TokenAnalysis = tokenAnalysis
	request.TextKey = textKey
	request.DefineModule = patternDefine
	return r.ValidFormatText(request)
}

func NewRepository() *Repository {
	var repository Repository
	repository.tokenAnalyser = NewTokenAnalyser()
	repository.patternKey = NewPatternKey()
	return &repository
}
