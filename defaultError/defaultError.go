package defaultError

import (
	"fmt"
	"os"
)

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func HandleDefaultError(returnValue interface{}, err error) interface{} {
	if (err != nil) {
		errorf("Something went wrong dude %+v", err)
	}
	return returnValue
}
