package errors

import "errors"

func IsTemporary(err error) bool {
	var t interface {
		IsTemporary() bool
	}
	return errors.As(err, &t) && t.IsTemporary()
}
