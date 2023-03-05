package rounds

import (
	"countdown/utils"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Letters struct{}

func (c *Letters) Generate() Result {
	// generate a pseudo-random seed
	rand.Seed(int64(timeNowF.YearDay()))

	// fetch a word from the seeds
	base := utils.LettersWords[rand.Intn(len(utils.LettersWords))]

	// buff out the word to 9 chars
	if len(base) != 9 {
		rem := 9 - len(base)
		base = base + randomString(rem)
	}

	baseRune := []rune(base)
	// shuffle characters randomly
	rand.Shuffle(len(base), func(i, j int) {
		baseRune[i], baseRune[j] = baseRune[j], baseRune[i]
	})

	today := strings.ToUpper(string(baseRune))

	// add the length of the word as a hint
	todayWithHint := fmt.Sprintf("%s (%d)", today, len(base))

	// re-seed to fetch yesterday's solution
	rand.Seed(int64(timeNowF.Add(-1 * time.Hour * 24).YearDay()))
	y := utils.LettersWords[rand.Intn(len(utils.LettersWords))]
	yesterday := strings.ToUpper(y)

	return Result{
		Round:     lettersRound,
		Today:     todayWithHint,
		Yesterday: yesterday,
	}
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) // a = 65 and z = 65 + 25
	}
	return string(bytes)
}
