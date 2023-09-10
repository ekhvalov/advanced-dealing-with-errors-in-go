package errors

import "errors"

func IsAny(err error, references ...error) bool {
	if nil == err {
		return false
	}
	for _, e := range references {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
