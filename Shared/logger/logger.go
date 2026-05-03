package logger

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {

	configLog := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := configLog.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
