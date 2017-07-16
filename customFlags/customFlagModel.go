package customFlags

type flags struct {
	debug bool
	logFile string
}

func (f flags) GetDebugMode() bool {
	return f.debug
}

func (f flags) GetLogFile() string {
	return f.logFile
}