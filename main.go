package main

import (
	"countdown/rounds"
	"countdown/sms"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var timeNowF = time.Now()

func main() {
	// generate a letters round
	l := rounds.Letters{}
	letters := l.Generate()

	// generate a numbers round
	// todo - this isn't yet implemented
	n := rounds.Numbers{}
	numbers := n.Generate()

	// generate the conundrum
	c := rounds.Conundrum{}
	conundrum := c.Generate()

	// generate a pseudo-random number
	// to determine whether today should
	// be a numbers or letters round
	rand.Seed(int64(timeNowF.YearDay()))
	today := rand.Int() % 2

	// re-seed and find yesterday
	rand.Seed(int64(timeNowF.Add(-1 * time.Hour * 24).YearDay()))
	yesterday := rand.Int() % 2

	// format today and yesterday's puzzles
	todayString := t(today, letters, conundrum, numbers)
	yesterdayString := y(yesterday, letters, conundrum, numbers)

	err := sms.Send(fmt.Sprintf("%s\n%s", todayString, yesterdayString))
	if err != nil {
		log.Fatal(err.Error())
	}
}

// todo - given the numbers round isn't finished
// we'll need to reintroduce switching logic for
// showing the numbers or letters round when it is
func t(today int, letters rounds.Result, conundrum rounds.Result, numbers rounds.Result) string {
	return rounds.TodayString(letters, conundrum)
}

func y(yesterday int, letters rounds.Result, conundrum rounds.Result, numbers rounds.Result) string {
	return rounds.YesterdayString(letters, conundrum)
}
