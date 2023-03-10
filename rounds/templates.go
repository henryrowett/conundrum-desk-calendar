package rounds

import "fmt"

const dateFormat = "02/01/2006"
const titleTemplate = "  \nš %s - HENRY'S COUNTDOWN CALENDAR - Day %d š”\n\n"
const todayTemplate = "Today's puzzles:\n%s puzzle: %s\n%s: %s\n\n"
const yesterdayTemplate = "Yesterday's answers:\n%s puzzle: %s\n%s: %s"

func Title() string {
	return fmt.Sprintf(titleTemplate, timeNowF.Format(dateFormat), timeNowF.YearDay())
}

func TodayString(puzzle Result, conundrum Result) string {
	return fmt.Sprintf(todayTemplate, puzzle.Round, puzzle.Today, conundrum.Round, conundrum.Today)
}

func YesterdayString(puzzle Result, conundrum Result) string {
	return fmt.Sprintf(yesterdayTemplate, puzzle.Round, puzzle.Yesterday, conundrum.Round, conundrum.Yesterday)
}
