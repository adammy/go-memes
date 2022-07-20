package httpapi

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig[Cfg comparable](path, env string, defaults map[string]interface{}) (*Cfg, error) {
	var config Cfg

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(env)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
