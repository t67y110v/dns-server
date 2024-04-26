package main

import (
	"github.com/t67y110v/dns-server/internal/app"
	"github.com/t67y110v/dns-server/internal/config"
	"github.com/t67y110v/dns-server/pkg/logger"
)

func main() {
	logger := logger.Get(logger.LevelInfo)
	logger.Info("logger init success")
	config, err := config.New(logger)
	if err != nil {
		return
	}
	logger.Info("configuration init success")
	app := app.New(config, logger)
	logger.Info("applicatioin init success")
	app.Start()

}
