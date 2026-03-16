package main

import "testing"

func TestCountFreq(t *testing.T) {
	got := countFreq([]string{"a", "b", "a", ""})
	if got["a"] != 2 || got["b"] != 1 {
		t.Fatalf("got=%v", got)
	}
	if _, ok := got[""]; ok {
		t.Fatalf("empty string should be ignored")
	}
}
