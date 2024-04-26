package models

type (
	Response struct {
		Endpoints      []Endpoint     `json:"endpoints"`
		VersionWeights map[string]int `json:"version_weights"`
		ServiceConfig  string         `json:"service_config"`
	}

	Endpoint struct {
		Address string `json:"address"`
		Version string `json:"version"`
		Weight  int    `json:"weight"`
	}
)
