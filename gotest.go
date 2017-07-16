package main

import (
	"fmt"
	"os"
	"errors"
	"customFlags"
	"./customLogger"
	saveRun "./defaultError"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func doubleInt(i int) (sum int, err error) {
	sum = i + i
	err = errors.New("Internal error in doubleInt")
	return
}

func main () {
	f := customFlags.Init()
	fmt.Println("DebugMode :" , f.GetDebugMode())
	fmt.Println("Logfile :" , f.GetLogFile())
	//logger1 := saveRun.HandleDefaultError(customLogger.Init("./log/log1.txt", "log1: "))
	logger3, _ := customLogger.Init("./log/log1.txt", "log3: ")
	logger2, _ := customLogger.Init("./log/log2.txt", "log2: ")
	//defer logger1.Close()
	defer logger3.Close()
	defer logger2.Close()

	logger3.Log("Gehts?")
	//logger1.Log("Logger geht")
	logger2.Log("Logger geht")
	double, err := doubleInt(3)
	//logger1.Log("Logger1 geht immernoch")
	logger2.Log("Logger2 geht immernoch")
	logger3.Log("Geht! Really!")
	if err != nil {
		errorf("Error in doubleInt: %+v ", err)
	}
	fmt.Println(double)
	//logger1.Log("Logger geht immernoch aber wegen Error nicht sichtbar")
}
