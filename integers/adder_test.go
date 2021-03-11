package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output: 5
}
