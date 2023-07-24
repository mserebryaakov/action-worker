package zaplog

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Default(debug bool, wr io.Writer) *zap.Logger {
	var level zapcore.Level

	if debug {
		level = zap.DebugLevel
	} else {
		level = zap.WarnLevel
	}

	var core zapcore.Core = zapcore.NewCore(getProductionEncoder(), zapcore.AddSync(wr), level)

	return zap.New(core)
}

func getProductionEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
