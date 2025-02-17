package main

import "testing"

func TestAdd(t *testing.T) {
	// given
	l, r := 1, 2
	expected := 3

	// when
	result := Add(l, r)

	// then
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", l, r, result, expected)
	}
}
