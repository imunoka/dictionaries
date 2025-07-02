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

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("non-existing word", func (t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("key")
		want := ErrNotFound

		assertError(t, err, want)
	})
}

func TestDictionary_Add(t *testing.T) {
	t.Run("new word", func (t *testing.T) {
		dictionary := Dictionary{}
		key := "key"
		value := "value"
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertValue(t, dictionary, key, value)
	})

	t.Run("existing word", func (t *testing.T) {
		key := "key"
		value := "value"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new value")

		assertError(t, err, ErrKeyExists)
		assertValue(t, dictionary, key, value)
	})
}

func TestDictionary_Update(t *testing.T) {
	t.Run("existing word", func (t *testing.T) {
		key := "key"
		value := "value"
		dictionary := Dictionary{key: value}
		newValue := "new value"
		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertValue(t, dictionary, key, newValue)
	})

	t.Run("non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("key", "new value")

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func ExampleDictionary_Search() {
	dictionary := Dictionary{"key": "value"}
	value, err := dictionary.Search("key")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("found:", value)

	// Output: found: value
}

func ExampleDictionary_Add() {
	dictionary := Dictionary{}
	err := dictionary.Add("key", "value")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	value, err := dictionary.Search("key")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("found:", value)

	// Output: found: value
}

func ExampleDictionary_Update() {
	dictionary := Dictionary{"key": "value"}
	err := dictionary.Update("key", "new value")

	if err != nil {
		fmt.Println("error:", err)
	}

	value, err := dictionary.Search("key")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("found:", value)

	// Ouput: found: new value
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertValue(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("error:", err)
	}
	assertStrings(t, got, value)
}
