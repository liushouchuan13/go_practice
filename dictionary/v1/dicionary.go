package v1

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) string {
	return dictionary[word]
}
