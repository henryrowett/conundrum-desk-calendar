package rounds

import "time"

const (
	lettersRound   = "letters"
	numbersRound   = "numbers"
	conundrumRound = "conundrum"
)

type Round interface {
	Generate() Result
}

type Result struct {
	Round     string
	Today     string
	Yesterday string
}

var TimeNowF = time.Now()
