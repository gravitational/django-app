package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/trace"
)

// InitLoggerCLI tools by default log into syslog, not stderr
func InitLoggerCLI() {
	log.SetLevel(log.InfoLevel)
	// clear existing hooks:
	log.StandardLogger().Hooks = make(log.LevelHooks)
	log.SetFormatter(&trace.TextFormatter{})
	log.SetOutput(os.Stderr)
}

// InitLoggerDebug configures the logger to dump everything to stderr
func InitLoggerDebug() {
	// clear existing hooks:
	log.StandardLogger().Hooks = make(log.LevelHooks)
	log.SetFormatter(&trace.TextFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}
