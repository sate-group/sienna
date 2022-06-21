package logger

import "log"

func Infof(format string, v ...any) {
	format = string(COLOR_GREEN) + "INFO: " + format + string(COLOR_RESET)
	log.Printf(format, v...)
}
