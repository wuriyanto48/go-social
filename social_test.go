package social

import (
	"testing"
)

func TestSocial(t *testing.T) {
	expected := "Google"

	googleType := Google

	if googleType.String() != expected {
		t.Error("Error Social type conversion")
	}
}
