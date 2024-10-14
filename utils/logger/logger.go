package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(name string) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}
	logger := zapLogger.Named(name)
	return logger, nil
}
