package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nikolas", "")
		want := "Hello, Nikolas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Nikolas", "Spanish")
		want := "Hola, Nikolas"
		assertCorrectMessage(t, got, want)
	})

	t.Run(`say "Hola, Mundo" when an empty string is supplied in Spanish`, func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Mundo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Nikolas", "French")
		want := "Bonjour, Nikolas"
		assertCorrectMessage(t, got, want)
	})

	t.Run(`say "Bonjour, Monde" when an empty string is supplied in French`, func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, Monde"
		assertCorrectMessage(t, got, want)
	})
}
