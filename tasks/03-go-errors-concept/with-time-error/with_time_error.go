package errs

import (
	"fmt"
	"time"
)

type WithTimeError struct { // Реализуй меня.
	err error
	t   time.Time
}

func (e *WithTimeError) Error() string {
	return fmt.Sprintf("%v, occurred at: %s", e.err, e.t)
}

func (e *WithTimeError) Time() time.Time {
	return e.t
}

func (e *WithTimeError) Unwrap() error {
	return e.err
}

func NewWithTimeError(err error) error {
	return newWithTimeError(err, time.Now)
}

func newWithTimeError(err error, timeFunc func() time.Time) error {
	return &WithTimeError{err: err, t: timeFunc()}
}
