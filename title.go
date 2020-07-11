package justwatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ContentPaths is a struct that contains the full paths of a Content
type ContentPaths struct {
	MovieDetailOverview  string `json:"MOVIE_DETAIL_OVERVIEW,omitempty"`
	SeasonDetailOverview string `json:"SEASON_DETAIL_OVERVIEW,omitempty"`
}

// Backdrop contains just a backdrop URL
type Backdrop struct {
	BackdropURL string `json:"backdrop_url"`
}

// Credit is a struct that contains information about roles involved in the
// content, like actors, music composers, producers, etc.
type Credit struct {
	Role          string `json:"role"`
	CharacterName string `json:"character_name"`
	PersonID      int    `json:"person_id"`
	Name          string `json:"name"`
}

// ExternalID is a struct that contains ids in other platforms
type ExternalID struct {
	Provider string `json:"provider"`
	ID       string `json:"external_id"`
}

// Content is a struct that contains details about a specific title.
type Content struct {
	JWEntityID          string            `json:"jw_entity_id"`
	ID                  int               `json:"id"`
	Title               string            `json:"title"`
	FullPath            string            `json:"full_path"`
	FullPaths           *ContentPaths     `json:"full_paths"`
	Poster              string            `json:"poster"`
	BackDrops           []*Backdrop       `json:"backdrops"`
	ShortDescription    string            `json:"short_description"`
	OriginalReleaseYear int               `json:"original_release_year"`
	TMDBPopularity      float32           `json:"tmdb_popularity"`
	ObjectType          string            `json:"object_type"`
	OriginalTitle       string            `json:"original_title"`
	Offers              []*Offer          `json:"offers"`
	Scoring             []*ContentScoring `json:"scoring"`
	Credits             []*Credit         `json:"credits"`
	ExternalIDs         []*ExternalID     `json:"external_ids"`
	GenreIDs            []int             `json:"genre_ids"`
	AgeCertification    string            `json:"age_certification"`
	Runtime             int               `json:"runtime"`
}

// GetContentByTypeAndID gets a content object from the JustWatch API. It accepts
// two parameters: a content type (movie, show, show_season, etc.) and an ID.
func (c *Client) GetContentByTypeAndID(contentType string, id int) (*Content, error) {
	var content Content
	endpoint := fmt.Sprintf("titles/%s/%d/locale/%s", contentType, id, c.locale)
	response, err := c.doReq(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(responseBody, &content); err != nil {
		return nil, err
	}
	return &content, nil
}
