package config

// Config is the entrypoint for all configuration items
type Config struct {
	Server Server `mapstructure:"server"`
	DSN    string `mapstructure:"dsn"`
	Log    Log    `mapstructure:"log"`
}

// Server configuration
type Server struct {
	HTTPListen  string    `mapstructure:"http_listen"`
	HTTPSListen string    `mapstructure:"https_listen"`
	TLS         ServerTLS `mapstructure:"tls"`
}

// ServerTLS is used for server TLS certificate configuration
type ServerTLS struct {
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
}

// Log configuration
type Log struct {
	Level string `mapstructure:"level"`
	Path  string `mapstructure:"path"`
}
