package config

type Config struct {
	BindAddr string `toml:"bind_addr"`
	BasePath string `toml:"base_path"`
	LogLevel string `toml:"log_level"`
}
