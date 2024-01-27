package parsinglogfiles

import (
	"regexp"
	"slices"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[[A-Z]{3}]`)
	validLogs := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}

	if err != nil {
		panic(err)
	}

	s := re.FindString(text)

	if slices.Contains(validLogs, s) {
		return true
	}

	return false
}

func SplitLogLine(text string) []string {
	panic("Please implement the SplitLogLine function")
}

func CountQuotedPasswords(lines []string) int {
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
