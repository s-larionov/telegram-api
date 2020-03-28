package models

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`

	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos"`
}
