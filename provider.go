package justwatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeeplinkPackages is a struct that contains packages coming from deeplink data
type DeeplinkPackages struct {
	AndroidTV string `json:"android_tv"`
	FireTV    string `json:"fire_tv"`
	TVOS      string `json:"tvos"`
}

// DeeplinkData is a struct that contains the deeplink data from a provider
type DeeplinkData struct {
	Scheme       string            `json:"scheme"`
	Packages     []string          `json:"packages"`
	Platforms    []string          `json:"platforms"`
	PathTemplate string            `json:"path_template"`
	Extras       map[string]string `json:"extras"`
}

// ProviderData is a struct that contains the data coming from a provider
type ProviderData struct {
	Deeplink []*DeeplinkData `json:"deeplink_data"`
}

// Provider is a struct that represents a provider from JustWatch API
type Provider struct {
	ID                int           `json:"id"`
	TechnicalName     string        `json:"technical_name"`
	ShortName         string        `json:"short_name"`
	ClearName         string        `json:"clear_name"`
	Priority          int           `json:"priority"`
	DisplayPriority   int           `json:"display_priority"`
	MonetizationTypes []string      `json:"monetization_types"`
	IconURL           string        `json:"icon_url"`
	Slug              string        `json:"slug"`
	Data              *ProviderData `json:"data"`
}

// GetProviders returns all the providers from the JustWatch API
func (c *Client) GetProviders() ([]*Provider, error) {
	var providers []*Provider
	endpoint := fmt.Sprintf("providers/locale/%s", c.locale)
	response, err := c.doReq(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(responseBody, &providers); err != nil {
		return nil, err
	}
	return providers, nil
}

// GetProviderByID returns a single provider that matches the id passed
// to the function
func (c *Client) GetProviderByID(id int) (*Provider, error) {
	providers, err := c.GetProviders()
	if err != nil {
		return nil, err
	}
	for _, provider := range providers {
		if id == provider.ID {
			return provider, nil
		}
	}
	return nil, fmt.Errorf("no providers that matches the id %d", id)
}
