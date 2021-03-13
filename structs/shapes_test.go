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
		name string
		s    Shape
		want float64
	}{
		{name: "Rectangle", s: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "Circle", s: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", s: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			got := tt.s.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.s, got, tt.want)
			}
		})
	}
}
