package errs

import (
	"errors"
	"strings"
)

// Pretty делает цепочку ошибок более читаемой – очередная заврапленная
// ошибка будет представлена на новой строке.
func Pretty(err error) error {
	if nil == err {
		return nil
	}
	return newPrettyError(err)
}

func newPrettyError(err error) *prettyError {
	pe := prettyError{err: err}
	sb := strings.Builder{}
	prevMsg := err.Error()
	sb.Grow(len(prevMsg))
	for err = errors.Unwrap(err); err != nil; err = errors.Unwrap(err) {
		msg := err.Error()
		sb.WriteString(strings.TrimSpace(prevMsg[:len(prevMsg)-len(msg)]))
		sb.WriteByte('\n')
		prevMsg = msg
	}
	sb.WriteString(prevMsg)
	pe.msg = sb.String()
	return &pe
}

type prettyError struct {
	msg string
	err error
}

func (e *prettyError) Unwrap() error {
	return e.err
}

func (e *prettyError) Error() string {
	return e.msg
}
