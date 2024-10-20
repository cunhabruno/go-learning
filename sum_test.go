package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(1, 4)
	if result != 5 {
		t.Errorf("error")
	}
}
