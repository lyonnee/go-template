package config

type AppConfig struct {
	Env         string `mapstructure:"env"`         // 环境变量
	Name        string `mapstructure:"name"`        // 应用名称
	Version     string `mapstructure:"version"`     // 应用版本
	Description string `mapstructure:"description"` // 应用描述
	HostId      int64  `mapstructure:"host_id"`     // 主机id
}
