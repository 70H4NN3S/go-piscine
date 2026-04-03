package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAllInOne(t *testing.T) {
	t.Run("single slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumAllInOne(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("multiple slices", func(t *testing.T) {
		num1 := []int{1, 2, 3}
		num2 := []int{4, 5}
		num3 := []int{7}

		got := SumAllInOne(num1, num2, num3)
		want := 22
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{123, 234, 345, 456, 654}
	b.ReportAllocs()
	for b.Loop() {
		Sum(numbers)
	}
}
