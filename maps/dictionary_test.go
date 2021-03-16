package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	t.Run("word exists", func(t *testing.T) {
		got, _ := dict.Search("test")
		assertStrings(t, got, "this is a test")
	})

	t.Run("work does not exist", func(t *testing.T) {
		_, err := dict.Search("missing")
		assertError(t, err, ErrKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	key, val := "test", "this is a test"
	dict.Add(key, val)
	got, err := dict.Search("test")
	if err != nil {
		t.Fatalf("key %q was not added", key)
	}
	assertStrings(t, got, "this is a test")
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
