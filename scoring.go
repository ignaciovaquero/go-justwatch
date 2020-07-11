package justwatch

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

// ContentScoring is a struct that contains scoring information for different
// providers.
type ContentScoring struct {
	ProviderType string  `json:"provider_type"`
	Value        float32 `json:"value"`
}
