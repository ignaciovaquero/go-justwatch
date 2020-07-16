package justwatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Genre is a struct that contains data for genres
type Genre struct {
	ID            int    `json:"id"`
	ShortName     string `json:"short_name"`
	TechnicalName string `json:"technical_name"`
	Translation   string `json:"translation"`
	Slug          string `json:"slug"`
}

// GetGenres gets all the genres in JustWatch
func (c *Client) GetGenres() ([]*Genre, error) {
	var genres []*Genre
	endpoint := fmt.Sprintf("genres/locale/%s", c.locale)
	response, err := c.doReq(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(responseBody, &genres); err != nil {
		return nil, err
	}
	return genres, nil
}

// GetGenreByID returns a single genre that matches the id passed
// to the function
func (c *Client) GetGenreByID(id int) (*Genre, error) {
	genres, err := c.GetGenres()
	if err != nil {
		return nil, err
	}
	for _, genre := range genres {
		if id == genre.ID {
			return genre, nil
		}
	}
	return nil, fmt.Errorf("no genres that matches the id %d", id)
}
