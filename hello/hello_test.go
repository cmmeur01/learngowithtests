package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello with a name (string) as input", func(t *testing.T) {
		got := Hello("chris", "")
		want := "Hi2u, chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world with empty string as input", func(t *testing.T) {
		got := Hello("", "")
		want := "Hi2u, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in espanol", func(t *testing.T) {
		got := Hello("chris", "Spanish")
		want := "Hola, chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("chris", "French")
		want := "Bonjour, chris"

		assertCorrectMessage(t, got, want)
	})

}
