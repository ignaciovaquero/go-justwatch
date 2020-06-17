package justwatch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type countryName struct {
	DE string `json:"de"`
	ES string `json:"es"`
	FR string `json:"fr"`
	HR string `json:"hr"`
	IT string `json:"it"`
	JA string `json:"JA"`
	NL string `json:"NL"`
	RU string `json:"RU"`
}

type localeResponse struct {
	ExposedURLPart   string      `json:"exposed_url_part"`
	IsStandardLocale bool        `json:"is_standard_locale"`
	FullLocale       string      `json:"full_locale"`
	I18NState        string      `json:"i18n_state"`
	ISO31662         string      `json:"iso_3166_2"`
	Country          string      `json:"country"`
	Currency         string      `json:"currency"`
	CurrencyName     string      `json:"currency_name"`
	CountryNames     countryName `json:"country_names"`
}

func (c *Client) getLocale() (string, error) {
	const defaultLocale = "es_ES"
	resp, err := c.doReq(http.MethodGet, "locales/state", nil)
	if err != nil {
		return "", err
	}
	response, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	var locales []*localeResponse
	err = json.Unmarshal(response, locales)
	if err != nil {
		return "", err
	}

	for _, locale := range locales {
		if locale.ISO31662 == c.Country || locale.Country == c.Country {
			return locale.FullLocale, nil
		}
	}
	return defaultLocale, nil
}
