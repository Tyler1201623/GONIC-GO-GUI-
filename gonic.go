// Package gonic provides a simple, beautiful, and pure Go GUI framework.
// It aims to be the PyQt for Go: easy to install, powerful, looks beautiful,
// and has a dead-simple file structure.
package gonic

import (
	"fmt"
	"log"
)

// initNativeRenderer initializes the native renderer
func initNativeRenderer() bool {
	return tryNativeRenderer()
}

// runNativeWindow runs a single window with the native renderer
func runNativeWindow(window *Window) {
	// Use the existing native app runner with a single window
	runNativeApp([]*Window{window})
}

// ShowAlert is an alias for Alert for backward compatibility
func ShowAlert(message string) {
	Alert(message)
}

// Additional utility functions that may be needed

// LogInfo logs an informational message
func LogInfo(message string) {
	log.Println("INFO:", message)
}

// LogError logs an error message
func LogError(err error) {
	if err != nil {
		log.Println("ERROR:", err)
	}
}

// PrintVersion prints the current version of the framework
func PrintVersion() {
	fmt.Println("Gonic Framework v0.1.0")
	fmt.Println("The PyQt for Go")
}
