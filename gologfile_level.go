package gologfile

import (
	"github.com/memUsins/golog"
	"go.uber.org/zap/zapcore"
)

func levelToZap(level golog.Level) zapcore.Level {
	switch level {
	case golog.DebugLevel:
		return zapcore.DebugLevel
	case golog.InfoLevel:
		return zapcore.InfoLevel
	case golog.WarnLevel:
		return zapcore.WarnLevel
	case golog.ErrorLevel:
		return zapcore.ErrorLevel
	case golog.FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
