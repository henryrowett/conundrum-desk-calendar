package main

import (
	"countdown/rounds"
	"countdown/utils"
	"fmt"
	"math/rand"
	"time"
)

var timeNowF = time.Now()

func main() {
	// generate a letters round
	l := rounds.Letters{}
	letters := l.Generate()

	// generate a numbers round
	n := rounds.Numbers{}
	numbers := n.Generate()

	// generate the conundrum
	c := rounds.Conundrum{}
	conundrum := c.Generate()

	// generate a pseudo-random seed
	rand.Seed(int64(timeNowF.YearDay()))
	// generate a random number and mod 2
	today := rand.Int() % 2

	// re-seed and find yesterday
	rand.Seed(int64(timeNowF.Add(-1 * time.Hour * 24).YearDay()))
	yesterday := rand.Int() % 2

	var todayString string
	switch today {
	// it's a letters round
	case 0:
		todayString = fmt.Sprintf(utils.TodayTemplate, letters.Round, letters.Today, conundrum.Round, conundrum.Today)
	// it's a numbers round
	case 1:
		todayString = fmt.Sprintf(utils.TodayTemplate, numbers.Round, numbers.Today, conundrum.Round, conundrum.Today)
	}

	var yesterdayString string
	switch yesterday {
	// it's a letters round
	case 0:
		yesterdayString = fmt.Sprintf(utils.YesterdayTemplate, letters.Round, letters.Yesterday, conundrum.Round, conundrum.Yesterday)
	// it's a numbers round
	case 1:
		yesterdayString = fmt.Sprintf(utils.YesterdayTemplate, numbers.Round, numbers.Yesterday, conundrum.Round, conundrum.Yesterday)
	}

	fmt.Println(todayString)
	fmt.Println(yesterdayString)
}
