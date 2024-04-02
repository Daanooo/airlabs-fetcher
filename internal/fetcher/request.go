package fetcher

import "github.com/gookit/config/v2"

type request struct {
	baseUrl  string
	apiKey   string
	endpoint string
	params   map[string]any
}

func NewRequest(endpoint string, params map[string]any) *request {
	return &request{
		baseUrl:  config.String("baseUrl"),
		apiKey:   config.String("apiKey"),
		endpoint: endpoint,
		params:   params,
	}
}
