package apm

import (
	"os"

	"dbm-services/common/go-pubpkg/logger"
)

// initLogger initialization log
func initLogger() {
	l := logger.New(os.Stdout, true, logger.InfoLevel, map[string]string{})
	logger.ResetDefault(l)
}

// init
func init() {
	initLogger()
}
