package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Harok")
	want := "Hello, Harok"

	if got != want {
		t.Errorf("got: %q want %q", got, want)
	}
}
