package logs

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

const (
	LevelConfigPath = "logger.level"
	PathConfigPath  = "logger.path"
)

var (
	ZapLogger   *zap.Logger
	OtLogger    *otelzap.Logger
	LoggerLevel = zap.NewAtomicLevel()
	Syncer      *ReopenableWriteSyncer
)
