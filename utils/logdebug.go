package utils

import "log"

// DebugLogging helper func to log debug
func DebugLogging(logstring string, enableDebug bool) {
	if enableDebug {
		log.Printf(logstring)
	}
}
