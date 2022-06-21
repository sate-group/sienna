package logger

import "log"

func Warnf(format string, v ...any) {
	format = string(COLOR_PURPLE) + "WARN: " + format + string(COLOR_RESET)
	log.Printf(format, v...)
}
