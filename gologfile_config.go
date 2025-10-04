package gologfile

import (
	"github.com/memUsins/golog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// FileConfig core config for adapter
type FileConfig struct {
	Enable bool

	Level    golog.Level
	LJLogger *lumberjack.Logger
}

// defaultFileConfig setting up default config
func defaultFileConfig() *FileConfig {
	return &FileConfig{
		Enable: true,
		Level:  golog.DebugLevel,
		LJLogger: &lumberjack.Logger{
			Filename:   "logs/logs.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     28,
			LocalTime:  false,
			Compress:   false,
		},
	}
}
