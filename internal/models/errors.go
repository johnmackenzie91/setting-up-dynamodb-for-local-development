package models

import "fmt"

type errMissingAttribute struct {
	Field string
	Struct string
}

func (e errMissingAttribute) Error() string {
	return fmt.Sprint("missing field: %s, on struct: %s", e.Field, e.Struct)
}
