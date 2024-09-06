package log

import (
	"fmt"
	"os"
)

func Log(err error) {
	fmt.Println(err)
}

func LogFatal(err error) {
	Log(err)
	os.Exit(1)
}

func CheckEndLogFatal(err error) {
	if err != nil {
		LogFatal(err)
	}
}
