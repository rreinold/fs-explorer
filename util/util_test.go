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

func TestSuccessIsForbiddenPath(t *testing.T) {
	input := [...]string{"foo", "foo/bar", "/", "foo/bar1"}
	for _, value := range input {
		out := IsForbiddenPath(value)
		if out {
			t.Errorf("Incorrectly forbidden clean path: %v", value)
		}
	}
}
