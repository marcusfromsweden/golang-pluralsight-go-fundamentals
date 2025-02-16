package main

import "testing"

func TestAdd(t *testing.T) {
	// given
	l, r := 1, 2
	expected := 3

	// when
	got := Add(l, r)

	// then
	if got != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", l, r, got, expected)
	}
}
