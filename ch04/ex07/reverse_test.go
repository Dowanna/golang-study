package ex07

import (
	"testing"
)

func TestRevUtf8(t *testing.T) {
	utf8String := []byte("あいうえお")
	expected := []byte("おえういあ")

	RevUtf8(utf8String)

	if !sliceEqual(utf8String, expected) {
		t.Errorf("error!\n got     : %s\n expected: %s", utf8String, expected)
	}
}

func sliceEqual(slice []byte, expected []byte) bool {
	for i, r := range slice {
		if r != expected[i] {
			return false
		}
	}
	return true
}
