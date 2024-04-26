package config

import (
	"flag"
	"log/slog"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func New(logger *slog.Logger) (*Config, error) {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
	flag.Parse()
	config := &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logger.Error("error while decoding config file ", err)
		return nil, err
	}
	return config, nil
}
