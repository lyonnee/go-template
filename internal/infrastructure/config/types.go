package config

import "time"

// ================== AppConfig ==================
type AppConfig struct {
	Env         string `yaml:"env"`         // 环境变量
	Name        string `yaml:"name"`        // 应用名称
	Version     string `yaml:"version"`     // 应用版本
	Description string `yaml:"description"` // 应用描述
	HostId      int64  `yaml:"host_id"`     // 主机id
}

// ================== DatabaseConfig ==================

// ================== HttpConfig ==================
// HttpConfig 包含 HTTP 服务的配置

type HttpConfig struct {
	Port string `yaml:"port"`
}

// ================== GRPCConfig ==================
// GRPCConfig 包含 gRPC 服务的配置
type GRPCConfig struct {
	Port string `yaml:"port"`
	// Max concurrent HTTP/2 streams per connection (per RPCs over a single connection)
	MaxConcurrentStreams uint32 `yaml:"max_concurrent_streams"`

	// Keepalive server parameters
	Keepalive struct {
		Time                  time.Duration `yaml:"time"`    // ping interval between pings
		Timeout               time.Duration `yaml:"timeout"` // ping ack timeout
		MaxConnectionIdle     time.Duration `yaml:"max_connection_idle"`
		MaxConnectionAge      time.Duration `yaml:"max_connection_age"`
		MaxConnectionAgeGrace time.Duration `yaml:"max_connection_age_grace"`
	} `yaml:"keepalive"`

	// Keepalive enforcement policy
	Enforcement struct {
		MinTime             time.Duration `yaml:"min_time"`
		PermitWithoutStream bool          `yaml:"permit_without_stream"`
	} `yaml:"enforcement"`

	// Interceptors toggles
	Interceptors struct {
		EnableLogging  bool `yaml:"enable_logging"`
		EnableRecovery bool `yaml:"enable_recovery"`
	} `yaml:"interceptors"`
}

// ==================  LogConfig ==================
