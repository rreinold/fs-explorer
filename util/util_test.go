package util

import (
	"testing"
)

func TestFailIsForbiddenPath(t *testing.T) {
	input := [...]string{"..", "~", "./../", "~/.ssh/id_rsa", "~/..", "/bar/.."}
	for _, value := range input {
		out := IsForbiddenPath(value)
		if !out {
			t.Errorf("failed to catch forbidden path: %v", value)
		}
	}
}
