package rounds

import (
	"countdown/utils"
	"math/rand"
	"strings"
	"time"
)

type Conundrum struct{}

func (c *Conundrum) Generate() Result {
	// generate a pseudo-random seed
	rand.Seed(int64(timeNowF.YearDay()))

	// fetch a word from the seeds
	base := utils.ConundrumWords[rand.Intn(len(utils.ConundrumWords))]

	// shuffle the characters randomly
	baseRune := []rune(base)
	rand.Shuffle(len(base), func(i, j int) {
		baseRune[i], baseRune[j] = baseRune[j], baseRune[i]
	})

	// format the string
	today := strings.ToUpper(string(baseRune))

	// re-seed to fetch yesterday's solution
	rand.Seed(int64(timeNowF.Add(-1 * time.Hour * 24).YearDay()))
	y := utils.ConundrumWords[rand.Intn(len(utils.ConundrumWords))]
	yesterday := strings.ToUpper(y)

	// return the result
	return Result{
		Round:     conundrumRound,
		Today:     today,
		Yesterday: yesterday,
	}
}
