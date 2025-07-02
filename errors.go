package dictionaries

const (
	ErrNotFound = DictionaryErr("could not find requested key/value pair")
	ErrKeyExists = DictionaryErr("cannot add key because it already exits")
	ErrKeyDoesNotExist = DictionaryErr("cannot perform operation on key because it does not exist")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}
