package config

import "time"

// Config 是应用的总配置结构体
type Config struct {
	AppName string        `mapstructure:"appName"`
	Server  ServerConfig  `mapstructure:"server"`
	Store   StoreConfig   `mapstructure:"store"`
	Tracing TracingConfig `mapstructure:"tracing"`
}

// ServerConfig 包含HTTP服务器相关的配置
type ServerConfig struct {
	Port            string        `mapstructure:"port"`
	LogLevel        string        `mapstructure:"logLevel"`
	LogFormat       string        `mapstructure:"json"` // "text" or "json"
	ReadTimeout     time.Duration `mapstructure:"readTimeout"`
	WriteTimeout    time.Duration `mapstructure:"writeTimeout"`
	IdleTimeout     time.Duration `mapstructure:"idleTimeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdownTimeout"`
}

// StoreConfig 包含与存储相关的配置
type StoreConfig struct {
	Type string `mapstructure:"type"`
	// DSN string `mapstructure:"dsn"` // Example for Postgres
}

// TracingConfig 包含与分布式追踪相关的配置
type TracingConfig struct {
	Enabled      bool    `mapstructure:"enabled"`
	OTELEndpoint string  `mapstructure:"otelEndpoint"`
	SampleRatio  float64 `mapstructure:"sampleRatio"`
}
