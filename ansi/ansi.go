package ansi

import (
	"errors"
	"regexp"
)

// https://github.com/acarl005/stripansi - thanks
const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}

// actualIndex returns the index in the original (ANSI included) string
// that corresponds to the given display index (ANSI excluded).
// Returns an error if the display index is out of range or if the ansiStr argument is an empty string.
func ActualIndex(ansiStr string, displayIndex int) (int, error) {

	actual := 0
	display := 0
	runes := []rune(ansiStr)

	if displayIndex < 0 || displayIndex >= len(runes) {
		return 0, errors.New("index out of range")
	}

	for i := 0; i < len(runes); {
		remaining := string(runes[i:])
		loc := re.FindStringIndex(remaining)

		if loc != nil && loc[0] == 0 {
			// Skip over ANSI escape sequence
			i += loc[1]
			actual += loc[1]
		} else {
			if display == displayIndex {
				return actual, nil
			}
			i++
			actual++
			display++
		}
	}

	return 0, errors.New("index out of range")
}

// ActiveANSICodes returns a slice of active ANSI codes at the given display index.
func ActiveANSICodes(s string, displayIndex int) []string {
	runes := []rune(s)
	activeCodes := []string{}
	display := 0

	for i := 0; i < len(runes); {
		remaining := string(runes[i:])
		loc := re.FindStringIndex(remaining)

		if loc != nil && loc[0] == 0 {
			code := string(runes[i : i+loc[1]])

			if code == "\x1b[0m" {
				activeCodes = []string{} // Reset
			} else {
				activeCodes = append(activeCodes, code)
			}
			i += loc[1]
		} else {
			if display == displayIndex {
				break
			}
			i++
			display++
		}
	}

	return activeCodes
}
