package ints

import "testing"

func TestSum(t *testing.T) {
	total := Sum(1, 2)
	expected := 3

	if total != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, total)
	}
}
