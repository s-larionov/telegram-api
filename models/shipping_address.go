package models

import (
	"fmt"
	"strings"
)

// ShippingAddress This object represents a shipping address.
type ShippingAddress struct {
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`

	// State, if applicable
	State string `json:"state,omitempty"`

	// City
	City string `json:"city"`

	// First line for the address
	StreetLine1 string `json:"street_line1"`

	// Second line for the address
	StreetLine2 string `json:"street_line2"`

	// Address post code
	PostCode string `json:"post_code"`
}

func (a ShippingAddress) String() string {
	address := fmt.Sprintf(
		"%s %s %s, %s\n%s\n%s",
		a.PostCode,
		a.CountryCode,
		a.State,
		a.City,
		a.StreetLine1,
		a.StreetLine2,
	)

	return strings.TrimSpace(address)
}
