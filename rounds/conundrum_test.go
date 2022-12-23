package rounds

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConundrumGenerate(t *testing.T) {
	// set the date so the pseudo-random
	// rand seed is consistent for testing
	timeNowF = time.Date(2020, 10, 10, 0, 0, 0, 0, time.UTC)

	// expected result
	today := "CALREFLYU"
	yesterday := "COLLEAGUE"

	// act
	round := Conundrum{}
	result := round.Generate()

	// assert
	assert.Equal(t, conundrumRound, result.Round)
	assert.Equal(t, today, result.Today)
	assert.Equal(t, yesterday, result.Yesterday)
}
