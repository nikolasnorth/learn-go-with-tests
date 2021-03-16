package maps

import "errors"

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("key not found")

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) {
	d[key] = val
}
