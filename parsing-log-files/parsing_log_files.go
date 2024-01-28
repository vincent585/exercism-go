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
	re, err := regexp.Compile(`<[~\-=*]*>`)

	if err != nil {
		panic(err)
	}

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`.*".*((?i)password(?-i)).*".*`)

	if err != nil {
		panic(err)
	}

	passwordCount := 0

	for _, line := range lines {
		if re.MatchString(line) {
			passwordCount++
		}
	}

	return passwordCount
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`((?i)end-of-line\w*(?-i))`)

	if err != nil {
		panic(err)
	}

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
