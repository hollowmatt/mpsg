package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet_en(t *testing.T) {
	got := greet("en")
	want := "Hello world"

	if got != want {
		//mark the test as failed
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreet_fr(t *testing.T) {
	got := greet("fr")
	want := "Bonjour le monde"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
