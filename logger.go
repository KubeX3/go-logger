package logger

import (
	"fmt"
	"time"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

// getTime returns the current time
func getTime() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// printLog handles the actual printing logic to avoid code duplication
func printLog(level string, color string, message string) {
	fmt.Printf("[%s] %s[%s]%s %s\n", getTime(), color, level, colorReset, message)
}

// ==========================================
// Standard Loggers (Handle lists of args)
// Usage: LogInfo("Server start", err)
// ==========================================

func LogInfo(args ...interface{}) {
	printLog("INFO", colorGreen, fmt.Sprint(args...))
}

func LogWarning(args ...interface{}) {
	printLog("WARNING", colorYellow, fmt.Sprint(args...))
}

func LogError(args ...interface{}) {
	printLog("ERROR", colorRed, fmt.Sprint(args...))
}

// ==========================================
// Formatted Loggers (Handle %v, %s, %d)
// Usage: LogInfoF("Server start port: %d", 8080)
// ==========================================

func LogInfoF(format string, args ...interface{}) {
	printLog("INFO", colorGreen, fmt.Sprintf(format, args...))
}

func LogWarningF(format string, args ...interface{}) {
	printLog("WARNING", colorYellow, fmt.Sprintf(format, args...))
}

func LogErrorF(format string, args ...interface{}) {
	printLog("ERROR", colorRed, fmt.Sprintf(format, args...))
}