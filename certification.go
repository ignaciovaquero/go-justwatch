package justwatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Certification is a struct that contains data for certifications
type Certification struct {
	ID            int    `json:"id"`
	TechnicalName string `json:"technical_name"`
	Description   string `json:"description"`
	ObjectType    string `json:"object_type"`
	Country       string `json:"country"`
	Order         int    `json:"order"`
	Organization  string `json:"organization"`
}

// GetCertifications gets all the certifications in JustWatch
func (c *Client) GetCertifications() ([]*Certification, error) {
	var certifications []*Certification
	endpoint := "age_certifications"
	response, err := c.doReq(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(responseBody, &certifications); err != nil {
		return nil, err
	}
	return certifications, nil
}

// GetCertificationByID returns a single certification that matches the id passed
// to the function
func (c *Client) GetCertificationByID(id int) (*Certification, error) {
	certifications, err := c.GetCertifications()
	if err != nil {
		return nil, err
	}
	for _, certification := range certifications {
		if id == certification.ID {
			return certification, nil
		}
	}
	return nil, fmt.Errorf("no certifications that matches the id %d", id)
}
