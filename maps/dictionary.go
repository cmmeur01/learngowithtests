package dictionary

const (
	errorNotFound   = KVMapErr("could not find the key you are looking for")
	errorKeyExists  = KVMapErr("key already exists")
	errorKeyMissing = KVMapErr("can not find the key to update OR delete")
)

type KVMapErr string

func (e KVMapErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", errorNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errorNotFound:
		d[key] = value
	case nil:
		return errorKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newVal string) error {
	_, err := d.Search(key)

	switch err {
	case errorNotFound:
		return errorKeyMissing
	case nil:
		d[key] = newVal
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

type Dictionary map[string]string
