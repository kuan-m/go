package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "test value"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "test value"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		want := "test value"
		word := "test"

		err := dict.Add(word, "test value")

		assertError(t, err, nil)
		got, err := dict.Search("word")
		if err != nil {
			t.Fatal("shoud find added word: ", err)
		}
		assertStrings(t, got, want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v want: %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v want: %v", got, want)
	}
}
