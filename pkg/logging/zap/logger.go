package zap

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nferreira/app/pkg/env"
	"github.com/nferreira/logging/pkg/logging"
)

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
	Panic = "panic"
	Fatal = "fatal"
)

type Logger struct {
	formatter logging.Formatter
	Log *zap.Logger
	logging.Organization
	logging.System
}

func New(formatter logging.Formatter) logging.Logger {
	return &Logger{
		formatter: formatter,
	}
}

func (l *Logger) CheckHealth(ctx context.Context) error {
	return nil
}

func (l *Logger) Start(ctx context.Context) error {
	cfg := zap.Config{
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,
		},
	}
	switch env.GetString("LOG_LEVEL", Debug) {
	case Debug:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case Info:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case Warn:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case Error:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case Panic:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case Fatal:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	}

	var err error
	l.Log, err = cfg.Build()
	if err != nil {
		return err
	}

	l.Organization = logging.Organization{
		Id:   env.GetString("ORG_ID", ""),
		Name: env.GetString("ORG_NAME", ""),
		Unit: env.GetString("ORG_UNIT", ""),
	}

	l.System = logging.System{
		Environment: env.GetString("ENV", "dev"),
		Id:          env.GetString("SYSTEM_ID", ""),
		Hostname:    l.hostname(),
		AppName:     env.GetString("APP_NAME", ""),
	}

	return nil
}

func (l *Logger) Stop(ctx context.Context) error {
	return nil
}

func (l *Logger) Fatalf(correlationId string, format string, args ...interface{}) {
	l.logf(l.Log.Fatal, correlationId, format, args...)
}

func (l *Logger) Fatal(correlationId string, message string) {
	l.log(l.Log.Fatal, correlationId, message)
}

func (l *Logger) Panicf(correlationId string, format string, args ...interface{}) {
	l.logf(l.Log.Panic, correlationId, format, args...)
}

func (l *Logger) Panic(correlationId string, message string) {
	l.log(l.Log.Panic, correlationId, message)
}

func (l *Logger) Errorf(correlationId string, format string, args ...interface{}) {
	l.logf(l.Log.Error, correlationId, format, args...)
}

func (l *Logger) Error(correlationId string, message string) {
	l.log(l.Log.Error, correlationId, message)
}

func (l *Logger) Infof(correlationId string, format string, args ...interface{}) {
	l.logf(l.Log.Info, correlationId, format, args...)
}

func (l *Logger) Info(correlationId string, message string) {
	l.log(l.Log.Info, correlationId, message)
}

func (l *Logger) Warnf(correlationId string, format string, args ...interface{}) {
	l.logf(l.Log.Warn, correlationId, format, args...)
}

func (l *Logger) Warn(correlationId string, message string) {
	l.log(l.Log.Warn, correlationId, message)
}

func (l *Logger) Debugf(correlationId string, format string, args ...interface{}) {
	l.logf(l.Log.Debug, correlationId, format, args...)
}

func (l *Logger) Debug(correlationId string, message string) {
	l.log(l.Log.Debug, correlationId, message)
}

func (l *Logger) logf(handler func(_ string, _ ...zap.Field),
	correlationId string,
	format string,
	args ...interface{}) {
	l.log(handler, correlationId, fmt.Sprintf(format, args...))
}

func (l *Logger) log(handler func(_ string, _ ...zap.Field),
	correlationId string,
	message string) {
	handler(l.formatter.Format(l.Organization, l.System, correlationId, message))
}

func (l *Logger) format(correlationId string, message string) string {
	return l.formatter.Format(l.Organization, l.System, correlationId, message)
}

func (l *Logger) hostname() string {
	if hostname, err := os.Hostname(); err != nil {
		return "unknown"
	} else {
		return hostname
	}
}
