package config

import (
	"fmt"
	"runtime/debug"
)

func ErrorLog(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	errLog.Output(2, trace)

}

func InfoLog(info string) {

	infoLog.Println(info)
}
