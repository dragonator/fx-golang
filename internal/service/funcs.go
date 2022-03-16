package service

import "go.uber.org/fx"

// RunFirst -
func RunFirst(logger ILogger) {
	logger.Log("FIRST call before application startup")
}

// RunSecond -
func RunSecond(logger ILogger) {
	logger.Log("SECOND call  before application startup")
}

// RunThird -
func RunThird(logger ILogger) {
	logger.Log("THIRD call  before application startup")
}

// ParamStruct -
type ParamStruct struct {
	fx.In

	Logger ILogger
}

// RunWithPAramStruct -
func RunWithPAramStruct(ps ParamStruct) {
	ps.Logger.Log("Using logger from parameter struct")
}
