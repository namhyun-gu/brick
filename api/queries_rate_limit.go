package api

import (
	"bytes"
)

type RateLimit struct {
	Limit     uint   `json:"limit"`
	Remaining uint   `json:"remaining"`
	Reset     uint64 `json:"reset"`
	Used      uint   `json:"used"`
}

type ResourcesRateLimit struct {
	Core                RateLimit `json:"core"`
	GraphQl             RateLimit `json:"graphql"`
	IntegrationManifest RateLimit `json:"integration_manifest"`
	Search              RateLimit `json:"search"`
}

type RateLimitStatus struct {
	Resources ResourcesRateLimit `json:"resources"`
	Rate      RateLimit          `json:"rate"`
}

func GetRateLimit(client *Client) (*RateLimitStatus, error) {
	var rateLimitStatus RateLimitStatus
	r := bytes.NewReader([]byte(`{}`))
	err := client.REST("https://api.github.com/", "GET", "rate_limit", r, &rateLimitStatus)
	if err != nil {
		return nil, err
	}
	return &rateLimitStatus, err
}
