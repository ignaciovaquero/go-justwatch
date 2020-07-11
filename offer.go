package justwatch

// OfferURLs defines a set of urls for the provider
type OfferURLs struct {
	StandardWeb string `json:"standard_web"`
}

// Offer defines an offer in JustWatch
type Offer struct {
	MonetizationType string     `json:"monetization_type"`
	ProviderID       int        `json:"provider_id"`
	RetailPrice      float32    `json:"retail_price,omitempty"`
	Currency         string     `json:"currency"`
	URLs             *OfferURLs `json:"urls"`
	PresentationType string     `json:"presentation_type"`
}
