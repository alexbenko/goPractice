package caveman

import "testing"

func TestCaveman(t *testing.T) {
	actual := Caveman("Javascript Sucks")
	expected := "jAvAsCrIpT sUcKs"

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
