package customLogger

func Init(filename string, prefix string) (customLogger myLogger, err error) {
	customLogger.init(filename, prefix)
	return
}
