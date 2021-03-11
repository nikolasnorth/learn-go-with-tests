package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := strings.Repeat("a", 6)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 6))
	// Output: aaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
