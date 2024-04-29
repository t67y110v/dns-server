package endpoints

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/t67y110v/dns-server/internal/models"
)

func UpdateEndpoints() {
	var resp = models.Service{
		Endpoints: []models.Endpoint{
			{
				Address: "127.0.0.1:8080",
				Version: "v1",
				Weight:  "35",
			},
			{
				Address: "127.0.0.1:8090",
				Version: "v2",
				Weight:  "34",
			},
			{
				Address: "127.0.0.1:8095",
				Version: "v2",
				Weight:  "67",
			},
		},
		VersionWeights: []models.VW{
			{
				Version: "v1",
				Weight:  "10",
			},
			{
				Version: "v2",
				Weight:  "90",
			},
		},
		ServiceConfig: "{\"loadBalancingPolicy\": \"habr_balancer\"}",
	}

	f, err := os.Create("configs/update_endpoints.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(resp)
	if err != nil {
		panic(err)
	}
}
