package v4

type Dictionary map[string]string

var (
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
