package rounds

import "time"

var timeNowF = time.Now()

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
