package parsinglogfiles

import (
	"regexp"
	"slices"
)

func IsValidLine(text string) bool {
	valid := []string{"TRC", "DBG", "INF", "WRN", "ERR", "FTL"}
	re := regexp.MustCompile(`^\[(.*)\]`)
	match := re.FindAllStringSubmatch(text, 1)

	if len(match) == 0 {
		return false
	}
	return slices.Contains(valid, match[0][1])
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*-=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*(?i)password.*"`)
	count := 0
	for _, line := range lines {
		match := re.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		match := re.FindAllStringSubmatch(line, 1)
		if len(match) > 0 {
			lines[i] = "[USR] " + match[0][1] + " " + lines[i]
		}
	}

	return lines
}
