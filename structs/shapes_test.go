package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := Perimeter(r.Width, r.Height)
	want := 40.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("Area for rectangle", func(t *testing.T) {
		r := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, r, want)
	})

	t.Run("Area for circle", func(t *testing.T) {
		c := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, c, want)
	})
}
