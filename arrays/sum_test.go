package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %q", got, want, numbers)
		}
	})
}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	fmt.Println(SumAll([]int{1, 2}, []int{0, 9}))
	// Output: [3 9]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum non-empty collections", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 5})
		want := []int{5, 14}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty collection", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 5})
		want := []int{0, 14}

		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	fmt.Println(SumAllTails([]int{1, 2, 3}, []int{0, 9, 5}))
	// Output: [5 14]
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails(SumAllTails([]int{1, 2, 3}, []int{0, 9, 5}))
	}
}
