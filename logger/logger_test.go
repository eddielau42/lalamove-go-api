package logger

import (
	"testing"
)

const (
	logfile = "lalamove.log"
)

func TestLogger(t *testing.T) {
	SetFile(logfile)
	SetLevel(INFO_LEVEL)

	Info("---> logger testing >> this is info output...")
	// Debug("---> logger testing >> this is debug output...")
	// Warn("---> logger testing >> this is warn output...")
	// Error("---> logger testing >> this is error output...")
}