package logs

import (
	"datalog-go/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var err error

func init() {
	utils.InitConfig()
	utils.InitTimeZone()

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(msg interface{}, fields ...zapcore.Field) {
	switch v := msg.(type) {
	case string:
		log.Info(v, fields...)
	}
}

func Debug(msg interface{}, fields ...zapcore.Field) {
	switch v := msg.(type) {
	case string:
		log.Debug(v, fields...)
	}
}

func Error(msg interface{}, fields ...zapcore.Field) {
	switch v := msg.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
