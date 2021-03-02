package util

import (
	"regexp"
)

func IsForbiddenPath(input string) bool {
	// Forbid '..' which could move out of hosted dir
	// Forbid '~' to prevent referencing user home dir
	REGEXP := `(\.\.|~)`
	match, err := regexp.Match(REGEXP, []byte(input))
	if err != nil {
		return true
	}
	return match
}
