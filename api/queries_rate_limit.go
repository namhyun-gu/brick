package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func GetRateLimit() (*RateLimitStatus, error) {
	url := "https://api.github.com/rate_limit"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed get rate limit (code: %d)", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rateLimitStatus RateLimitStatus
	err = json.Unmarshal(body, &rateLimitStatus)
	if err != nil {
		return nil, err
	}
	return &rateLimitStatus, err
}
