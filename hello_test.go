package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Anthony")
	want := "Hello, Anthony"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
