package customLogger

import (
	"os"
	"log"
)

type myLogger struct {
	file   *os.File
	logger *log.Logger
}

func (l myLogger) Log(text string) {
	l.logger.Println(text)
}

func (l myLogger) Close() {
	l.file.Close()
}

func (l *myLogger) init(filename string, prefix string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //os.O_WRONLY|os.O_CREATE|
	if err != nil {
		log.Fatal(err)
	}
	l.file = f
	l.logger = log.New(f, prefix, 5)
}
