package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// Convert the string to mutable slice of runes.
	runes := []rune(log)
	// Modify the slice.
	for i, r := range runes {
		if r == oldRune {
			runes[i] = newRune
		}
	}
	// Convert back to string.
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    return utf8.RuneCountInString(log) <= limit
}
