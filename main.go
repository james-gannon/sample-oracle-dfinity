package main

import (
	"time"

	framework "github.com/hyplabs/dfinity-oracle-framework"
	"github.com/hyplabs/dfinity-oracle-framework/models"
)

func main() {
	tokyoEndpoints := []models.Endpoint{
		{
			Endpoint: "http://api.weatherapi.com/v1/current.json?key=1943150c046041f9b29172550220111&q=Tokyo,JP",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.current.temp_c",
			},
		},
		{
			Endpoint: "https://api.weatherbit.io/v2.0/current?key=9e5d89305a0a4e8b8fb381ad3a1c1627&city=Tokyo&country=JP",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.data[0].temp",
			},
		},
	}
	delhiEndpoints := []models.Endpoint{
		{
			Endpoint: "http://api.weatherapi.com/v1/current.json?key=1943150c046041f9b29172550220111&q=Delhi,IN",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.current.temp_c",
			},
		},
		{
			Endpoint: "https://api.weatherbit.io/v2.0/current?key=9e5d89305a0a4e8b8fb381ad3a1c1627&city=Delhi&country=IN",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.data[0].temp",
			},
		},
	}
	config := models.Config{
		CanisterName:   "sample_oracle",
		UpdateInterval: 5 * time.Minute,
	}
	engine := models.Engine{
		Metadata: []models.MappingMetadata{
			{Key: "Tokyo", Endpoints: tokyoEndpoints},
			{Key: "Delhi", Endpoints: delhiEndpoints},
		},
	}
	oracle := framework.NewOracle(&config, &engine)
	oracle.Bootstrap()
	oracle.Run()
}