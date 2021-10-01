package models

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

// PollType Poll type, currently can be “regular” or “quiz”
type PollType string
