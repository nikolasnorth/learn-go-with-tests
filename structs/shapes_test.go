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

	areaTests := []struct {
		s    Shape
		want float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.s.Area()
		if got != tt.want {
			t.Errorf("got %f want %f", got, tt.want)
		}
	}
}
