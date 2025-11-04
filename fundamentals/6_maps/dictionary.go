package maps

import "errors"

var (
	ErrNotFound   = errors.New("could not find the word")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word, value string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
