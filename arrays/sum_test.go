package arrays

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func ExampleSum() {
	fmt.Println(Sum([5]int{1, 2, 3, 4, 5}))
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([5]int{1, 2, 3, 4, 5})
	}
}
