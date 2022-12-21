package rounds

import "fmt"

const todayTemplate = "Today's puzzles:\n\n%s puzzle:\n%s\n\n%s:\n%s\n\n"
const yesterdayTemplate = "Yesterday's answers:\n\n%s puzzle:\n%s\n\n%s:\n%s"

func TodayString(puzzle Result, conundrum Result) string {
	return fmt.Sprintf(todayTemplate, puzzle.Round, puzzle.Today, conundrum.Round, conundrum.Today)
}

func YesterdayString(puzzle Result, conundrum Result) string {
	return fmt.Sprintf(yesterdayTemplate, puzzle.Round, puzzle.Yesterday, conundrum.Round, conundrum.Yesterday)
}
