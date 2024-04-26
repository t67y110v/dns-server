package endpoints

import "github.com/t67y110v/dns-server/internal/models"

// TODO: parse some configuration file to fill structure fields
func GetResponse() any {
	var resp = models.Response{
		Endpoints: []models.Endpoint{
			{
				Address: "127.0.0.1:8080",
				Version: "v1",
			},
			{
				Address: "127.0.0.1:8090",
				Version: "v2",
				Weight:  34,
			},
			{
				Address: "127.0.0.1:8095",
				Version: "v2",
				Weight:  67,
			},
		},
		VersionWeights: map[string]int{
			"v1": 10,
			"v2": 90,
		},
		ServiceConfig: "{\"loadBalancingPolicy\": \"habr_balancer\"}",
	}

	return resp
}
