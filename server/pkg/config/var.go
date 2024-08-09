package config

var (
	envParam = &EnvParam{}
	config   = &Config{}
)

func GetEnv() *EnvParam {
	return envParam
}

func GetConfig() *Config {
	return config
}
