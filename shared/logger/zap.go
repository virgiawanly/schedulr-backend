package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type ZapLogger struct {
	Logger *zap.Logger
}

func NewZapLogger(isDevelopment bool, options ...zap.Option) (Logger, error) {
	var logger *zap.Logger
	var err error

	if isDevelopment {
		logger, err = zap.NewDevelopment(options...)
		if err != nil {
			return nil, err
		}
	} else {
		logger, err = zap.NewProduction(options...)
		if err != nil {
			return nil, err
		}
	}

	return &ZapLogger{
		Logger: logger,
	}, nil
}

func (z *ZapLogger) Debug(msg string) {
	z.Logger.Debug(msg)
}

func (z *ZapLogger) Debugf(format string, args ...any) {
	z.Logger.Debug(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) Info(msg string) {
	z.Logger.Info(msg)
}

func (z *ZapLogger) Infof(format string, args ...any) {
	z.Logger.Info(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) Warn(msg string) {
	z.Logger.Warn(msg)
}

func (z *ZapLogger) Warnf(format string, args ...any) {
	z.Logger.Warn(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) Error(msg string) {
	z.Logger.Error(msg)
}

func (z *ZapLogger) Errorf(format string, args ...any) {
	z.Logger.Error(fmt.Sprintf(format, args...))
}

func (z *ZapLogger) Fatal(msg string) {
	z.Logger.Fatal(msg)
}

func (z *ZapLogger) Fatalf(format string, args ...any) {
	z.Logger.Fatal(fmt.Sprintf(format, args...))
}
