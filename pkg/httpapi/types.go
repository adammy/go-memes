package httpapi

import (
	"time"
)

// ServerType denotes the implementation of the Server interface.
type ServerType string

// LogLevel denotes the log level to use (e.g. ["trace", "debug", "info", "warn", "error", "critical"]).
type LogLevel string

// Error is a structured error to display on HTTP services.
type Error struct {
	Error     string    `json:"error"`
	Timestamp time.Time `json:"time"`
}

// ServerConfig defines configuration for the Server.
type ServerConfig struct {
	// Port defines the port the server runs on.
	Port int `mapstructure:"port"`

	// Type defines the implementation of Server.
	Type ServerType `mapstructure:"type"`

	// Logger defines configuration for the HTTP logger.
	Logger LoggerConfig `mapstructure:"logger"`
}

// LoggerConfig defines configuration for an HTTP logger.
type LoggerConfig struct {
	// JSON defines if the log should be printed in structured json or pretty printed as lines.
	JSON bool `mapstructure:"json"`

	// AppName defines the name of the app.
	AppName string `mapstructure:"app_name"`

	// AppVersion defines the version of the app.
	AppVersion string `mapstructure:"app_version"`

	// EnvName defines the environment the app is running in.
	EnvName string `mapstructure:"env_name"`

	// Level denotes the LogLevel to use for this logger.
	Level LogLevel `mapstructure:"level"`
}
