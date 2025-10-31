package config

import (
	"regexp"

	"github.com/spf13/viper"
)

type Config struct {
	ContainerName   string `mapstructure:"container_name"`
	Pattern         string `mapstructure:"pattern"`
	Template        []byte `mapstructure:"template"`
	ActionName      string `mapstructure:"action_name"`
	CompiledPattern *regexp.Regexp
}

func GetConfigFromViper(viper_instance *viper.Viper) (*Config, error) {
	app_config := Config{}
	viper_instance.BindEnv("pattern")
	viper_instance.BindEnv("template")
	viper_instance.BindEnv("action_name")
	viper_instance.BindEnv("container_name")
	viper_instance.AutomaticEnv()
	err := viper_instance.UnmarshalExact(&app_config)
	if err != nil {
		return nil, err
	}
	compiled_pattern, err := regexp.Compile(app_config.Pattern)
	if err != nil {
		return nil, err
	}
	app_config.CompiledPattern = compiled_pattern
	return &app_config, nil
}
