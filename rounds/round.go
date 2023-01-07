package rounds

import "time"

var timeNowF = time.Now()

const dateFormat = "02-01-2006"

const (
	lettersRound   = "letter"
	numbersRound   = "number"
	conundrumRound = "conundrum"
)

type round interface {
	Generate() Result
}

type Result struct {
	Round     string
	Today     string
	Yesterday string
}
