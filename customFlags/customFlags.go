package customFlags

import (
	"flag"
)

func registerStringFlag(text *string, short string, long string, preset string, description string) {
	flag.StringVar(text, short, preset, description)
	flag.StringVar(text, long, preset, description)
}

func registerBoolFlag(value *bool, short string, long string, preset bool, description string) {
	flag.BoolVar(value, short, preset, description)
	flag.BoolVar(value, long, preset, description)
}

func Init() (customFlags flags) {
	registerStringFlag(&customFlags.logFile, "l", "log", "", "logFile")
	registerBoolFlag(&customFlags.debug, "d", "debug", false, "debugMode")
	setFlag(flag.CommandLine)
	flag.Parse()
	return
}
