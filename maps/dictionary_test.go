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
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		key, val := "test", "this is a test"

		err := dict.Add(key, val)
		if err != nil {
			t.Fatalf("got unexpected error %q", err)
		}

		got, err := dict.Search("test")
		if err != nil {
			t.Fatalf("key %q was not added", key)
		}
		assertStrings(t, got, "this is a test")
	})

	t.Run("existing word", func(t *testing.T) {
		key, val := "test", "this is a test"
		dict := Dictionary{key: val}
		err := dict.Add(key, val)
		assertError(t, err, ErrKeyExists)
	})
}

func TestUpdate(t *testing.T) {
	key, val := "test", "this is a test"
	dict := Dictionary{key: val}

	t.Run("key exists", func(t *testing.T) {
		newVal := "updated message"
		err := dict.Update(key, newVal)
		if err != nil {
			t.Fatalf("got unexpected error %q", err)
		}

		got, err := dict.Search(key)
		if err != nil {
			t.Fatalf("key %q was not updated", key)
		}
		assertStrings(t, got, newVal)
	})

	t.Run("key does not exist", func(t *testing.T) {
		err := dict.Update("missingKey", "this key is missing")
		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	key, val := "test", "this is a test"
	dict := Dictionary{key: val}
	dict.Delete(key)
	_, err := dict.Search(key)
	assertError(t, err, ErrKeyNotFound)
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
