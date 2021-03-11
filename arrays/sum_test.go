package arrays

import (
	"fmt"
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
