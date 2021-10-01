package models

// Venue This object represents a venue.
type Venue struct {
	// Venue location
	Location *Location `json:"location"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue.
	// For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.
	FoursquareType string `json:"foursquare_type,omitempty"`
}
