package logger

import "log"

type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...any)
	Info(msg string)
	Infof(format string, args ...any)
	Warn(msg string)
	Warnf(format string, args ...any)
	Error(msg string)
	Errorf(format string, args ...any)
	Fatal(msg string)
	Fatalf(format string, args ...any)
}

var defaultLogger Logger

func init() {
	zapLogger, err := NewZapLogger(true)
	if err != nil {
		log.Panicf("failed to initialize zap logger: %v", err)
	}
	SetGlobal(zapLogger)
}

func SetGlobal(logger Logger) {
	defaultLogger = logger
}

func GetGlobal() Logger {
	return defaultLogger
}

func Debug(msg string) {
	defaultLogger.Debug(msg)
}

func Debugf(format string, args ...any) {
	defaultLogger.Debugf(format, args...)
}

func Info(msg string) {
	defaultLogger.Info(msg)
}

func Infof(format string, args ...any) {
	defaultLogger.Infof(format, args...)
}

func Warn(msg string) {
	defaultLogger.Warn(msg)
}

func Warnf(format string, args ...any) {
	defaultLogger.Warnf(format, args...)
}

func Error(msg string) {
	defaultLogger.Error(msg)
}

func Errorf(format string, args ...any) {
	defaultLogger.Errorf(format, args...)
}

func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

func Fatalf(format string, args ...any) {
	defaultLogger.Fatalf(format, args...)
}
