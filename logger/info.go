package logger

import "log"

const (
	COLOR_RESET = "\033[0m"
	COLOR_GREEN = "\033[32m"
)

func Infof(format string, v ...any) {
	format = string(COLOR_GREEN) + "INFO: " + format + string(COLOR_RESET)
	log.Printf(format, v...)
}
