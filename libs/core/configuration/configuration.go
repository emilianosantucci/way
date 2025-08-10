package configuration

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func NewConfiguration() (config Configuration, err error) {
	config = Configuration{
		Profile: "Development",
		DB:      DatabaseConfiguration{},
		Web:     WebConfiguration{Host: "localhost", Port: 3000},
	}

	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err = viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return
		}
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	viper.WatchConfig()

	fmt.Printf("Configuration: %+v\n", config)

	return
}
