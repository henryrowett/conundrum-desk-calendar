package rounds

import (
	"fmt"
	"time"
)

const titleTemplate = "  \n🕛 HENRY'S COUNTDOWN CALENDAR %s 🕡\n\n"
const todayTemplate = "Today's puzzles:\n%s puzzle: %s\n%s: %s\n\n"
const yesterdayTemplate = "Yesterday's answers:\n%s puzzle: %s\n%s: %s"

func Title() string {
	return fmt.Sprintf(titleTemplate, time.Now().Format("02/01/2006"))
}

func TodayString(puzzle Result, conundrum Result) string {
	return fmt.Sprintf(todayTemplate, puzzle.Round, puzzle.Today, conundrum.Round, conundrum.Today)
}

func YesterdayString(puzzle Result, conundrum Result) string {
	return fmt.Sprintf(yesterdayTemplate, puzzle.Round, puzzle.Yesterday, conundrum.Round, conundrum.Yesterday)
}
