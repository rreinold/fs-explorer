package util

import (
	"regexp"
)

func IsForbiddenPath(input string) bool {
	REGEXP := `\.\.`
	match, err := regexp.Match(REGEXP, []byte(input))
	if err != nil {
		return true
	}
	return match
}
