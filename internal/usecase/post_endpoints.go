package usecase

import (
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/t67y110v/dns-server/internal/models"
)

func SetEndpoints(service string, entity models.Service) error {
	services := &models.AllServices{}
	_, err := toml.DecodeFile("configs/endpoints.toml", &services)
	if err != nil {
		slog.Info("error while decoding config file ", err)
		return err
	}

	services.Entity = append(services.Entity, entity)

	f, err := os.Create("configs/endpoints.toml")
	if err != nil {
		return err
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(services)
	if err != nil {
		return err
	}
	return nil
}
