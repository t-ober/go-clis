package add

import "testing"

func TestAdd(t *testing.T) {
	a, b := 1, 2
	res := add(a, b)
	exp := 3
	if res != exp {
		t.Fatalf("Expected %d, got %d", res, exp)
	}
}
