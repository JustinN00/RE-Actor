package config

type Config struct {
	ContainerName  string `mapstructure:"container_name"`
	Pattern    string `mapstructure:"pattern"`
	Template   string `mapstructure:"template"`
	ActionName string `mapstructure:"action_name"`
}
