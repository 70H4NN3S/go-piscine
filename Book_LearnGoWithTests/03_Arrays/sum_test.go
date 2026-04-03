package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{123, 234, 345, 456, 654}
	b.ReportAllocs()
	for b.Loop() {
		Sum(numbers)
	}
}
