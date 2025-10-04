package gologfile

import (
	"fmt"
	"github.com/memUsins/golog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// fileAdapter implements FileAdapter
type fileAdapter struct {
	cfg    *FileConfig
	writer *zap.Logger
}

// Log to file
func (a *fileAdapter) Log(log golog.Log) {
	if !a.cfg.Enable || !a.cfg.Level.IsEnabled(log.Level) {
		return
	}

	a.Format(&log)

	var fields []zap.Field
	if log.Data.Error != nil {
		fields = append(fields, zap.String("error", log.Data.Error.Error()))
	}

	if len(log.Data.Fields) > 0 {
		for key, value := range log.Data.Fields {
			fields = append(fields, zap.Any(key, value))
		}
	}

	switch log.Level {
	case golog.DebugLevel:
		a.writer.Debug(log.Message, fields...)
	case golog.InfoLevel:
		a.writer.Info(log.Message, fields...)
	case golog.WarnLevel:
		a.writer.Warn(log.Message, fields...)
	case golog.ErrorLevel:
		a.writer.Error(log.Message, fields...)
	case golog.FatalLevel:
		a.writer.Fatal(log.Message, fields...)
	case golog.UnselectedLevel:
		a.writer.Info(log.Message, fields...)
	default:
		a.writer.Info(log.Message, fields...)
	}
}

// Format formatting input log
func (a *fileAdapter) Format(log *golog.Log) {
	if log.Data.Name != "" {
		log.Data.Name = fmt.Sprintf("[%s]: ", log.Data.Name)
		log.Message = log.Data.Name + log.Message
	}
}

// NewFileAdapter returns new FileAdapter
func NewFileAdapter(cfg *FileConfig) FileAdapter {
	return &fileAdapter{
		cfg:    cfg,
		writer: newWriter(cfg),
	}
}

// NewDefaultFileAdapter returns new FileAdapter with default config
func NewDefaultFileAdapter() FileAdapter {
	cfg := defaultFileConfig()

	return &fileAdapter{
		cfg:    cfg,
		writer: newWriter(cfg),
	}
}

// NewDefaultFileAdapterWithLevel returns new FileAdapter with default config and target level
func NewDefaultFileAdapterWithLevel(level golog.Level) FileAdapter {
	cfg := defaultFileConfig()
	cfg.Level = level

	return &fileAdapter{
		cfg:    cfg,
		writer: newWriter(cfg),
	}
}

// newWriter returns new *zap.Logger
func newWriter(cfg *FileConfig) *zap.Logger {
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(pe)

	ioWriter := cfg.LJLogger
	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(ioWriter), levelToZap(cfg.Level))
	return zap.New(core)
}
