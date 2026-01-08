package config

import "time"

// ================== AppConfig ==================
type AppConfig struct {
	Env         string `mapstructure:"env"`         // 环境变量
	Name        string `mapstructure:"name"`        // 应用名称
	Version     string `mapstructure:"version"`     // 应用版本
	Description string `mapstructure:"description"` // 应用描述
	HostId      int64  `mapstructure:"host_id"`     // 主机id
}

// ================== DatabaseConfig ==================

// ================== HttpConfig ==================
// HttpConfig 包含 HTTP 服务的配置

type HttpConfig struct {
	Port string `mapstructure:"port"`
}

// ================== GRPCConfig ==================
// GRPCConfig 包含 gRPC 服务的配置
type GRPCConfig struct {
	Port string `mapstructure:"port"`
	// Max concurrent HTTP/2 streams per connection (per RPCs over a single connection)
	MaxConcurrentStreams uint32 `mapstructure:"max_concurrent_streams"`

	// Keepalive server parameters
	Keepalive struct {
		Time                  time.Duration `mapstructure:"time"`    // ping interval between pings
		Timeout               time.Duration `mapstructure:"timeout"` // ping ack timeout
		MaxConnectionIdle     time.Duration `mapstructure:"max_connection_idle"`
		MaxConnectionAge      time.Duration `mapstructure:"max_connection_age"`
		MaxConnectionAgeGrace time.Duration `mapstructure:"max_connection_age_grace"`
	} `mapstructure:"keepalive"`

	// Keepalive enforcement policy
	Enforcement struct {
		MinTime             time.Duration `mapstructure:"min_time"`
		PermitWithoutStream bool          `mapstructure:"permit_without_stream"`
	} `mapstructure:"enforcement"`

	// Interceptors toggles
	Interceptors struct {
		EnableLogging  bool `mapstructure:"enable_logging"`
		EnableRecovery bool `mapstructure:"enable_recovery"`
	} `mapstructure:"interceptors"`
}

// ==================  LogConfig ==================
