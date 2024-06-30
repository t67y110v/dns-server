package usecase

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/t67y110v/dns-server/internal/models"
)

// TODO: parse some configuration file to fill structure fields
func GetServiceConfiguration(service string) any {
	services := &models.AllServices{}
	_, err := toml.DecodeFile("configs/endpoints.toml", &services)
	if err != nil {
		log.Println("error while decoding config file ", err)
		return nil
	}

	serviceConfig := &models.Service{}

	for _, r := range services.Entity {
		if r.ServiceName == service {
			serviceConfig.ServiceName = r.ServiceName
			serviceConfig.Endpoints = r.Endpoints
			serviceConfig.VersionWeights = r.VersionWeights
			serviceConfig.ServiceConfig = r.ServiceConfig
		}
	}
	return serviceConfig
}
