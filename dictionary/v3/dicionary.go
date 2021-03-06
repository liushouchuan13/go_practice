package v3

import "errors"

type Dictionary map[string]string

var (
	ErrWordExists = errors.New("cannot add word because it already exists")
	ErrNotFound   = errors.New("could not find the word you were looking for")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
