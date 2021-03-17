package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrKeyExists   = DictionaryErr("key already exists")
	ErrKeyNotFound = DictionaryErr("key not found")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

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
	case nil:
		return ErrKeyExists
	case ErrKeyNotFound:
		d[key] = val
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}

	d[key] = val
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
