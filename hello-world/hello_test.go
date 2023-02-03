package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sergio")
	want := "Hello, Sergio"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}