package logger

import "go.uber.org/zap"

func New() (*zap.Logger, error) {
	// Production config:
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	
	return cfg.Build()
}