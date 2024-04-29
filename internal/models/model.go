package models

type AllServices struct {
	Entity []Service `json:"service" toml:"service"`
}
type Service struct {
	Endpoints      []Endpoint `json:"endpoints" toml:"endpoints"`
	VersionWeights []VW       `json:"version_weights" toml:"version_weights"`
	ServiceConfig  string     `json:"service_config" toml:"service_configs"`
	ServiceName    string     `json:"service_name" toml:"service_name"`
}

type Endpoint struct {
	Address string `json:"address" toml:"address"`
	Version string `json:"version" toml:"version"`
	Weight  string `json:"weight" toml:"weight"`
}

type VW struct {
	Version string `toml:"version"`
	Weight  string `toml:"weight"`
}
