package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Anthony", "")
		want := "Hello, Anthony"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Defaulting to saying hello to the world if name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Anthony", "spanish")
		want := "Hola, Anthony"
		assertCorrectMessage(t, got, want)
	})

}
