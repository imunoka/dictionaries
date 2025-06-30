package dictionaries

const (
	ErrNotFound = DictionaryErr("could not find requested key")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}
