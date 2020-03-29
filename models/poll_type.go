package models

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

// Poll type, currently can be “regular” or “quiz”
type PollType string
