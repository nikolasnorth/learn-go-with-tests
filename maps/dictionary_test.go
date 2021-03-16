package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	t.Run("word exists", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("work does not exist", func(t *testing.T) {
		_, err := dict.Search("missing")
		assertError(t, err, ErrKeyNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
