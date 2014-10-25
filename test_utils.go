package dockercommand

import (
	"bytes"
	"testing"
)

func expectEqual(t *testing.T, value, expected string) {
	if value != expected {
		t.Errorf("Expected value is: `%s`, Real value is `%s`", expected, value)
	}
}

func expectIsIn(t *testing.T, slice []string, expect string) {
	for _, value := range slice {
		if value == expect {
			return
		}
	}
	t.Errorf("`%s` is not present in slice `%s`", expect, slice)
}

func expectBytesEqual(t *testing.T, value, expected []byte) {
	if bytes.Compare(value, expected) != 0 {
		t.Errorf("Expected value is: `%s`, Real value is `%s`", expected, value)
	}
}
