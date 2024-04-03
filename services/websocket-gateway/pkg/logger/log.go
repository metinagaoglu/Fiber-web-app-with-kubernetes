package logger

import (
	"github.com/fatih/color"
)

func Info(module string, log string) {
	green := color.New(color.FgGreen).PrintfFunc()
	green("[INFO] [%s] ", module)

	notice := color.New(color.Bold, color.FgWhite).PrintlnFunc()
	notice(log)
}

func Error(module string, log string) {
	red := color.New(color.FgRed).PrintfFunc()
	red("[ERROR] [%s] ", module)

	notice := color.New(color.Bold, color.FgWhite).PrintlnFunc()
	notice(log)
}
