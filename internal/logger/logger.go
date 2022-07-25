package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[AntiEffort] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	// TODO: add log rotation
}

// Info prints log message with INFO level
func Info(format string, args ...any) {
	logger.Printf("[INFO] "+format+"\n", args...)
}

// Warn prints log message with WARN level
func Warn(format string, args ...any) {
	logger.Printf("[WARN] "+format+"\n", args...)
}

// Error prints log message with ERROR level
func Error(format string, args ...any) {
	logger.Printf("[ERROR] "+format+"\n", args...)
}

// Fatal prints log message with FATAL level and calls os.Exit(1)
func Fatal(format string, args ...any) {
	logger.Printf("[FATAL] "+format+"\n", args...)
	os.Exit(1)
}
