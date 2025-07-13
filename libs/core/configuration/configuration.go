package configuration

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewConfiguration() Configuration {
	return Configuration{
		Profile: "Development",
		DB:      DatabaseConfiguration{},
		Web:     WebConfiguration{Host: "localhost", Port: 3000},
	}
}

type ReadResult struct {
	fx.Out
	Configuration
}

func Read() (ReadResult, error) {
	config := NewConfiguration()

	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return ReadResult{}, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return ReadResult{}, err
	}

	viper.WatchConfig()

	fmt.Printf("Configuration: %+v\n", config)

	return ReadResult{Configuration: config}, nil
}
