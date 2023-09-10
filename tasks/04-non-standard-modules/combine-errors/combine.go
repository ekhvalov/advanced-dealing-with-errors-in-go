package errors

import (
	"fmt"
)

// Combine "прицепляет" ошибки other к err так, что они начинают фигурировать при выводе
// её на экран через спецификатор `%+v`. Если err является nil, то первостепенной ошибкой
// становится первая из ошибок other.
func Combine(err error, other ...error) error {
	var i int
	for ; err == nil && i < len(other); i++ {
		if other[i] != nil {
			err = other[i]
		}
	}
	if nil == err {
		return nil
	}
	ce := &combinedError{main: err}
	ce.extra = make([]error, 0, len(other))
	for ; i < len(other); i++ {
		if other[i] != nil {
			ce.extra = append(ce.extra, other[i])
		}
	}
	return ce
}

type combinedError struct {
	main  error
	extra []error
}

func (ce *combinedError) Unwrap() error {
	return ce.main
}

func (ce *combinedError) Error() string {
	return ce.main.Error()
}

func (ce *combinedError) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		_, _ = fmt.Fprint(state, ce.main.Error())
		if state.Flag('+') && len(ce.extra) > 0 {
			fmt.Fprintln(state)
			for _, err := range ce.extra {
				fmt.Fprintf(state, "  - %s\n", err.Error())
			}
		}
	}
}
