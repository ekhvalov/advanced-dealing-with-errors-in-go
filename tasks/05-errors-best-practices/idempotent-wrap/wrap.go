package errors

import "github.com/pkg/errors"

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Wrap оборачивает ошибку в сообщение. Также добавляет к ошибке стектрейс,
// если в цепочке уже нет ошибки со стектрейсом.
func Wrap(err error, msg string) error {
	var stErr stackTracer
	if errors.As(err, &stErr) {
		return errors.WithMessage(err, msg)
	}
	return errors.Wrap(err, msg)
}
