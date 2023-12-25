package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	applications := map[rune]string{'❗': "recommendation", '🔍': "search", '☀': "weather"}

	for _, c := range log {
		if val, ok := applications[c]; ok {
			return val
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	out := strings.Map(func(r rune) rune {
		if r == oldRune {
			return newRune
		}
		return r
	}, log)
	return out
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
