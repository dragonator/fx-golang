package service

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
