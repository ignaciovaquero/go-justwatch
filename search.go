package justwatch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SearchQuery is used for searching items in JustWatch. It contains a full
// list of filters.
type SearchQuery struct {
	AgeCertifications  []string          `json:"age_certifications,omitempty"`
	Genres             []string          `json:"genres,omitempty"`
	Languages          string            `json:"languages,omitempty"`
	MinPrice           int               `json:"min_price,omitempty"`
	MatchingOffersOnly bool              `json:"matching_offers_only,omitempty"`
	MaxPrice           int               `json:"max_price,omitempty"`
	MonetizationTypes  []string          `json:"monetization_types,omitempty"`
	PresentationTypes  []string          `json:"presentation_types,omitempty"`
	ReleaseYearFrom    int               `json:"release_year_from,omitempty"`
	ReleaseYearUntil   int               `json:"release_year_until,omitempty"`
	ScoringFilterTypes ScoringFilterType `json:"scoring_filter_types,omitempty"`
	TimelineType       string            `json:"timeline_type,omitempty"`
	SortBy             string            `json:"sort_by,omitempty"`
	SortAsc            bool              `json:"sort_asc,omitempty"`
	Providers          []string          `json:"providers,omitempty"`
	ContentTypes       []string          `json:"content_types,omitempty"`
	Page               int               `json:"page,omitempty"`
	PageSize           int               `json:"page_size,omitempty"`
}

// Item is used for returning a response from a search
type Item struct {
	ID         int      `json:"id"`
	JWEntityID string   `json:"jw_entity_id"`
	Title      string   `json:"title"`
	FullPath   string   `json:"full_path"`
	Poster     string   `json:"poster"`
	ObjectType string   `json:"object_type"`
	Offers     []*Offer `json:"offers"`
}

// SearchProvider is a response from a provider in a search API call
type SearchProvider struct {
	Total      int     `json:"total"`
	ProviderID int     `json:"provider_id"`
	Items      []*Item `json:"items"`
}

// SearchDay corresponds to a day in a response from the search API call
type SearchDay struct {
	Date      string            `json:"date"`
	Providers []*SearchProvider `json:"providers"`
}

// SearchResponse wraps a response from the search API call
type SearchResponse struct {
	SkipCount int          `json:"skip_count"`
	Days      []*SearchDay `json:"days"`
	Page      int          `json:"page"`
	PageSize  int          `json:"page_size"`
}

func (i *Item) String() string {
	return fmt.Sprintf("Item of type %s: %s", i.ObjectType, i.Title)
}

func (c *Client) search(kind string, query *SearchQuery) (*SearchResponse, error) {
	var resp SearchResponse
	endpoint := fmt.Sprintf("titles/%s/%s", c.locale, kind)
	body, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	response, err := c.doReq(http.MethodPost, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(responseBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchNew searches for new items. It can receive a SearchQuery in order to
// filter the results to be shown.
func (c *Client) SearchNew(query *SearchQuery) (*SearchResponse, error) {
	return c.search("new", query)
}

// SearchPopular searches for popular items. It can receive a SearchQuery in order to
// filter the results to be shown.
func (c *Client) SearchPopular(query *SearchQuery) (*SearchResponse, error) {
	return c.search("popular", query)
}
