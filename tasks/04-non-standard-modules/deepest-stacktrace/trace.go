package trace

import (
	"errors"
	"github.com/getsentry/sentry-go"
)

// GetDeepestStackTrace достаёт самый глубокий стектрейс из цепочки ошибок.
func GetDeepestStackTrace(err error) *sentry.Stacktrace {
	var stack *sentry.Stacktrace
	for err != nil {
		if st := sentry.ExtractStacktrace(err); st != nil {
			stack = st
		}
		err = errors.Unwrap(err)
	}
	return stack
}
