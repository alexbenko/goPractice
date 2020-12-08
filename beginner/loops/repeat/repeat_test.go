package repeat

import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}
