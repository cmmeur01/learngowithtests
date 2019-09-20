package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("notatest")

		assertError(t, got, errorNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("taiwan", "numba one")

		assertError(t, err, nil)
		assertKey(t, dictionary, "taiwan", "numba one")
	})

	t.Run("existing word", func(t *testing.T) {
		key := "poop"
		value := "stinky"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "so big")

		assertError(t, err, errorKeyExists)
		assertKey(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "one"
		value := "1"
		dictionary := Dictionary{key: value}
		newVal := "ein"

		dictionary.Update(key, newVal)

		assertKey(t, dictionary, key, newVal)
	})

	t.Run("new key", func(t *testing.T) {
		key := "one"
		value := "1"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, errorKeyMissing)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "one"
		value := "1"
		dictionary := Dictionary{key: value}

		dictionary.Delete(key)

		_, err := dictionary.Search(key)

		if err != errorNotFound {
			t.Errorf("expected %q to be deleted", key)
		}
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func assertKey(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should have found key after adding", err)
	}

	if got != value {
		t.Errorf("got %q wanted %q", got, value)
	}
}
