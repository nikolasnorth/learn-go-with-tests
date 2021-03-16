package maps

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("key not found")
var ErrKeyExists = errors.New("key already exists")

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotFound:
		d[key] = val
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}
