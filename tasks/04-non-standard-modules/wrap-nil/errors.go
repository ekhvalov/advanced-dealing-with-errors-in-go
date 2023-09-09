package errors

import "fmt"

// Wrapf работает аналогично fmt.Errorf, только поддерживает nil-ошибки.
func Wrapf(err error, f string, v ...any) error {
	if nil == err {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(f, v...), err)
}
