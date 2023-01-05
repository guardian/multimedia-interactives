package common

import (
	"testing"
)

func TestGenerateNumericId(t *testing.T) {

	got := GenerateNumericId()

	// Verify the result
	if got > 9999999 {
		t.Errorf("GenerateNumericId() = %d, want less than 9999999", got)
	}
	if got < 9900000 {
		t.Errorf("GenerateNumericId() = %d, want more than 9900000", got)
	}
}
