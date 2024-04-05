package fetcher

import "github.com/gookit/config/v2"

type request struct {
	baseUrl  string
	apiKey   string
	endpoint string
	params   []string
}

func NewRequest(endpoint string, params []string) *request {
	return &request{
		baseUrl:  config.String("baseUrl"),
		apiKey:   config.String("apiKey"),
		endpoint: endpoint,
		params:   params,
	}
}
