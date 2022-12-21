package rounds

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLettersGenerate(t *testing.T) {
	// rand seed is consistent for testing
	// set the date so the pseudo-random
	// set up
	TimeNowF = time.Date(2020, 10, 10, 0, 0, 0, 0, time.UTC)

	// expected result
	today := "mcefloaeo"
	yesterday := "generate"

	// act
	round := Letters{}
	result := round.Generate()

	// assert
	assert.Equal(t, lettersRound, result.Round)
	assert.Equal(t, today, result.Today)
	assert.Equal(t, yesterday, result.Yesterday)
}
