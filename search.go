package justwatch

import (
	"fmt"
	"net/http"
)

// ScoringFilter contains an integer value that allows to set a scoring filter
type ScoringFilter struct {
	MinScoringValue int `json:"min_scoring_value"`
}

// ScoringFilterType is a struct that gathers scoring filters from IMDB and
// Rotten Tomatoes
type ScoringFilterType struct {
	TomatoMeter ScoringFilter `json:"tomato:meter,omitempty"`
	IMDB        ScoringFilter `json:"imdb:score,omitempty"`
}

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

func (c *Client) searchItem(kind string) {
	endpoint := fmt.Sprintf("titles/%s/%s", c.locale, kind)
	c.doReq(http.MethodPost, endpoint, body)
}
