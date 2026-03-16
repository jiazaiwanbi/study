package main

import "testing"

func TestDedupSort(t *testing.T) {
	got := dedupSort([]int{3, 1, 2, 2, 3, 0})
	want := []int{0, 1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("len=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%d want=%d", i, got[i], want[i])
		}
	}
}
