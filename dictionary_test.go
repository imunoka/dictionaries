package dictionaries

import (
	"fmt"
	"testing"
)

func TestDictionary_Search(t *testing.T) {
	t.Run("existing word", func (t *testing.T) {
		dictionary := Dictionary{"key": "value"}
		got, err := dictionary.Search("key")

		want := "value"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("non-existing word", func (t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("key")

		want := ErrNotFound

		assertError(t, err, want)
	})
}

func ExampleDictionary_Search() {
	dictionary := Dictionary{"key": "value"}
	value, err := dictionary.Search("key")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Found:", value)

	// Output: Found: value
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but did not expect one")
	}
}
