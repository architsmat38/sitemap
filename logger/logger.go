package logger

import (
	"log"
)

/**
 * Prints INFO logs
 */
func Info(args ...interface{}) {
	log.Printf("[INFO] %v\n", args)
}

/**
 * Prints DEBUG logs
 */
func Debug(args ...interface{}) {
	log.Printf("[DEBUG] %v\n", args)
}

/**
 * Prints ERROR logs
 */
func Error(args ...interface{}) {
	log.Printf("[ERROR] %v\n", args)
}

/**
 * Prints custom logs
 */
func Print(args ...interface{}) {
	log.Print(args...)
}
