package config

import (
	"regexp"

	"github.com/spf13/viper"
)

type Config struct {
	ContainerName     string `mapstructure:"container_name"`
	Pattern           string `mapstructure:"pattern"`
	Template          string `mapstructure:"template"`
	ActionName        string `mapstructure:"action_name"`
	DiscordWebHookURL string `mapstructure:"discord_webhook_url"`
	LogLevel          string `mapstructure:"log_level"`
	CompiledPattern   *regexp.Regexp
}

func GetConfigFromViper(viper_instance *viper.Viper) (*Config, error) {
	app_config := Config{}
	viper_instance.BindEnv("pattern")
	viper_instance.BindEnv("template")
	viper_instance.BindEnv("action_name")
	viper_instance.SetDefault("action_name", "discord_webhook")
	viper_instance.BindEnv("container_name")
	viper_instance.BindEnv("discord_webhook_url")
	viper_instance.BindEnv("log_level")
	viper_instance.SetDefault("log_level", "INFO")
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
