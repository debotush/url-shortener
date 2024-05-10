package config

// Config represents the configuration for the application.
type Config struct {
	ServerAddress string // ServerAddress is the address on which the server will listen.
}

// LoadConfig loads the application configuration with default values.
func LoadConfig() *Config {
	return &Config{
		ServerAddress: ":8080", // Default server address is ":8080".
	}
}
