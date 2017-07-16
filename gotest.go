package main

import (
	"fmt"
	"os"
	"errors"
	"customFlags"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func doubleInt(i int) (sum int, err error) {
	sum = i + i
	err = errors.New("Error in doubleInt")
	return
}

func main () {
	f := customFlags.Init()
	fmt.Println("DebugMode :" , f.GetDebugMode())
	fmt.Println("Logfile :" , f.GetLogFile())

	double, err := doubleInt(3)
	if (err != nil) {
		errorf("Error in doubleInt: ", err)
	}
	fmt.Println(double)
}
