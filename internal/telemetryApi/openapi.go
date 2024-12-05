package telemetryApi

import (
	// Need a blank import to use go:embed.
	_ "embed"
	"errors"
	"gopkg.in/yaml.v3"
	"net/url"
)

//go:embed openapi.json
var Schema []byte

type openAPIServer struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

type openAPI struct {
	Servers []openAPIServer `json:"servers"`
}

// GetLocalDevHost returns Development Host from the OpenAPI specs.
func GetLocalDevHost() (string, error) {
	var s openAPI
	if err := yaml.Unmarshal(Schema, &s); err != nil {
		return "", err
	}

	for _, server := range s.Servers {
		if server.Description == "Local Development" {
			localURL, err := url.Parse(server.URL)
			if err != nil {
				return "", err
			}
			return localURL.Host, nil
		}
	}

	return "", errors.New("no servers found")
}
