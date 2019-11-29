package main

import "testing"

func TestMain(t *testing.T) {
	x := add(5, 3)
	if x != int(8) {
		t.Errorf("Wrong output. Expected: 8 but got: %v", x)
	}
}
