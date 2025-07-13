package configuration

import (
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"strings"
)

type LoggerParams struct {
	fx.In
	Configuration
}

type LoggerResult struct {
	fx.Out
	*zap.Logger
	*zap.SugaredLogger
}

func NewLogger(params LoggerParams) LoggerResult {
	var logger *zap.Logger
	var err error

	switch strings.ToLower(params.Configuration.Profile) {
	case "development", "dev":
		fmt.Printf("Logger config to: 'dev' mode.\n")
		logger, err = zap.NewDevelopment()
	case "production", "prod":
		fmt.Printf("Logger config to: 'prod' mode.\n")
		logger, err = zap.NewProduction()
	default:
		fmt.Printf("Unknown environment type: %s.\n", params.Configuration.Profile)
		logger = zap.NewExample()
	}

	if err != nil {
		log.Fatalf("Could not initialize zap logger: %v", err)
	}

	return LoggerResult{Logger: logger, SugaredLogger: logger.Sugar()}
}
