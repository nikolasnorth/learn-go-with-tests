package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is a test"}
	got := Search(dict, "test")
	want := "this is a test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
