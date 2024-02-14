package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Error(msg interface{}, fields ...zapcore.Field) {
	switch v := msg.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
