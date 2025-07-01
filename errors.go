package dictionaries

const (
	ErrNotFound = DictionaryErr("could not find requested key/value pair")
	ErrKeyExists = DictionaryErr("cannot add key because it already exits")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}
