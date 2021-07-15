// expected.go is a set of testing functions
package springytools

import (
	"testing"
)

func expectedInt(t *testing.T, expected int, got int) {
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func expectedBytes(t *testing.T, expected []byte, got []byte) {
	if len(expected) != len(got) {
		t.Errorf("expected same length %d for %q, got %d for %q", len(expected), expected, len(got), got)
	}
	for i, b := range expected {
		if (i < len(got)) && (got[i] != b) {
			t.Errorf("expected (diff at %d) %q, got %q", i, expected, got)
			break
		}
	}
}

func expectedString(t *testing.T, expected string, got string) {
	if expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
