package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("new error")
	errIsExist    = errors.New("already exist word")
	errIsNotExist = errors.New("word not exist in dictionary")
)

// Search for value in Dictionary
func (d Dictionary) Search(word string) (string, error) {

	str, exist := d[word]
	if exist {
		return str, nil
	}
	return "", errNotFound
}

// Add key & value in Dictionary
func (d Dictionary) Add(key, word string) error {

	_, err := d.Search(key)
	switch err {
	case errNotFound:
		d[key] = word
	case nil:
		return errIsExist
	}
	return nil
}

// Update key,word pair in Dictionary
func (d Dictionary) Update(key, word string) error {
	_, err := d.Search(key)
	switch err {
	case errNotFound:
		return errIsNotExist
	case nil:
		d[key] = word

	}
	return nil
}

// Delete word
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
