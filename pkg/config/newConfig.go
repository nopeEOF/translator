package config

func NewConfig(configPath string) (Config, error) {
	return getConfig(configPath)
}