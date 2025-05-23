package config

func Http() HttpConfig {
	return conf.Http
}

type HttpConfig struct {
	Port string `mapstructure:"port"`
}
