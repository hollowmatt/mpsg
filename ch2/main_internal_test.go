package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	got := greet()
	want := "Hello world"

	if got != want {
		//mark the test as failed
		t.Errorf("got %q want %q", got, want)
	}
}
