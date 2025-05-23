package config

func Cache() CacheConfig {
	return conf.Cache
}

type CacheConfig struct {
	Redis RedisConfig `mapstructure:"redis"`
}

type RedisConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Database  int    `mapstructure:"database"`
	Framework string `mapstructure:"framework"`
	Prefix    string `mapstructure:"prefix"`
}

func (conf RedisConfig) IsCluster() bool {
	return conf.Framework == "cluster"
}
