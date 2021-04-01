package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	loggerImpl  bsLogger
)

type BSLogger interface {
}

type bsLogger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	zapLogger, err := logConfig.Build()
	if err != nil {
		log.Fatalln(err)
	}

	loggerImpl = bsLogger{log: zapLogger}
}

/// Info prints an info level log with message
func Info(message string, tag ...zap.Field) {
	loggerImpl.log.Info(message, tag...)
	loggerImpl.log.Sync()
}

/// Error prints an error level log with message and error
func Error(message string, err error, tag ...zap.Field) {
	zapErr := zap.NamedError("error", err)
	tag = append(tag, zapErr)
	loggerImpl.log.Error(message, tag...)
	loggerImpl.log.Sync()
}
