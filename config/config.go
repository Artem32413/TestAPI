package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func configZap() zap.Config {
	return zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}
}

func ZapFunc() *zap.Logger {
	config := configZap()

	logger, err := config.Build()
	if err != nil{
		logger.Error(err.Error())
		return nil
	}

	defer logger.Sync()
	
	return logger
}