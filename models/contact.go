package models

// Contact This object represents a phone contact.
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Contact's user identifier in Telegram
	UserID int64 `json:"user_id,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard
	VCard string `json:"vcard,omitempty"`
}
