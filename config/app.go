package config

type AppConfig struct {
	Env         string `mapstructure:"env"`         // 环境变量
	Name        string `mapstructure:"name"`        // 应用名称
	Version     string `mapstructure:"version"`     // 应用版本
	Description string `mapstructure:"description"` // 应用描述
	NodeId      int64  `mapstructure:"node_id"`     // 节点ID，用于分布式ID生成
}
