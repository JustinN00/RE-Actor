package config

type Config struct {
	Pattern    string `mapstructure:"pattern"`
	Template   string `mapstructure:"template"`
	ActionName string `mapstructure:"action_name"`
}
