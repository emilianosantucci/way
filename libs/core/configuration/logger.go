package configuration

import (
	"fmt"
	"log"
	"strings"

	"go.uber.org/zap"
)

func NewLogger(config Configuration) (logger *zap.Logger, sugarLogger *zap.SugaredLogger, err error) {
	switch strings.ToLower(config.Profile) {
	case "development", "dev":
		fmt.Printf("Logger config to: 'dev' mode.\n")
		logger, err = zap.NewDevelopment()
	case "production", "prod":
		fmt.Printf("Logger config to: 'prod' mode.\n")
		logger, err = zap.NewProduction()
	default:
		fmt.Printf("Unknown environment type: %s.\n", config.Profile)
		logger = zap.NewExample()
	}

	if err != nil {
		log.Fatalf("Could not initialize zap logger: %v", err)
	}

	sugarLogger = logger.Sugar()

	return
}
