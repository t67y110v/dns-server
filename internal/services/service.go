package services

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/t67y110v/dns-server/internal/models"
)

func GetServices() []string {
	services := &models.AllServices{}
	_, err := toml.DecodeFile("configs/endpoints.toml", &services)
	if err != nil {
		log.Println("error while decoding config file ", err)
		return nil
	}

	names := make([]string, 0)

	for _, r := range services.Entity {
		names = append(names, r.ServiceName)
	}

	return names
}
