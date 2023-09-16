package rest

import "fmt"

func Handle() error {
	if err := usefulWork(); err != nil {
		return fmt.Errorf("%w: %v", ErrInternalServerError, err)
	}
	return nil
}

var usefulWork = func() error {
	return nil
}
