package main

import (
	"countdown/rounds"
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

	todayString := t(today, letters, conundrum, numbers)
	yesterdayString := y(yesterday, letters, conundrum, numbers)

	fmt.Println(fmt.Sprintf("%s\n%s", todayString, yesterdayString))

	fmt.Println(numbers)
}

func t(today int, letters rounds.Result, conundrum rounds.Result, numbers rounds.Result) string {
	switch today {
	// it's a letters round
	case 0:
		return rounds.TodayString(letters, conundrum)
	// it's a numbers round
	case 1:
		return rounds.TodayString(numbers, conundrum)
	}
	return ""
}

func y(yesterday int, letters rounds.Result, conundrum rounds.Result, numbers rounds.Result) string {
	switch yesterday {
	// it's a letters round
	case 0:
		return rounds.YesterdayString(letters, conundrum)
	// it's a numbers round
	case 1:
		return rounds.YesterdayString(numbers, conundrum)
	}
	return ""
}
