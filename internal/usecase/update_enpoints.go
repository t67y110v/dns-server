package usecase

import (
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/t67y110v/dns-server/internal/models"
)

func UpdateEndpoints(service string, entity models.Service) error {
	services := &models.AllServices{}
	_, err := toml.DecodeFile("configs/endpoints.toml", &services)
	if err != nil {
		slog.Info("error while decoding config file ", err)
		return err
	}
	for i, r := range services.Entity {
		if r.ServiceName == service {
			services.Entity[i].ServiceName = entity.ServiceName
			services.Entity[i].Endpoints = entity.Endpoints
			services.Entity[i].VersionWeights = entity.VersionWeights
			services.Entity[i].ServiceConfig = entity.ServiceConfig
		}
	}
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
