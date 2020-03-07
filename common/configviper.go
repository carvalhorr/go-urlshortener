package common

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func ConfigViper(env string, propertiesPath ...string) error {

	if env == "" {
		return fmt.Errorf("env cannot be empty")
	}

	if len(propertiesPath) < 1 {
		return fmt.Errorf("property path passed cannot be empty")
	}

	for _, path := range propertiesPath {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("properties")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("unable to read base configuration file. error = %+v", err)
	}

	envFileName := strings.Join([]string{"properties", ".", env}, "")

	viper.SetConfigName(envFileName)
	err = viper.MergeInConfig()
	if err != nil {
		return fmt.Errorf("configuration for the environment = %s not found. error = %+v", env, err)
	}
	return nil
}
