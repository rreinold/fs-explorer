package util

import (
	"testing"
)

func TestFailIsForbiddenPath(t *testing.T) {
	input := [...]string{"..", "~"}
	for _, value := range input {
		out := IsForbiddenPath(value)
		if !out {
			t.Errorf("failed to catch forbidden path: %v", value)
		}
	}
}
