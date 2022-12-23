package main

import (
	"countdown/rounds"
	"countdown/sms"
	"fmt"
	"log"
)

func main() {
	// generate a letters round
	l := rounds.Letters{}
	letters := l.Generate()

	// generate the conundrum
	c := rounds.Conundrum{}
	conundrum := c.Generate()

	// todo - we need to select whether to
	// use the letters or numbers round.
	// check the git history for this file.
	// for the moment we just send the
	// letters and conundrum round every day.

	// format today and yesterday's puzzles
	body := fmt.Sprintf("%s%s\n%s", rounds.Title(), rounds.TodayString(letters, conundrum), rounds.YesterdayString(letters, conundrum))
	err := sms.Send(body)
	if err != nil {
		log.Fatal(err.Error())
	}
}
