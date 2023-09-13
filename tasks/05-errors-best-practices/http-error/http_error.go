package httperr

import (
	"fmt"
	"net/http"
)

// Реализуй нас.
var (
	ErrStatusOK                  error = HTTPError(http.StatusOK)
	ErrStatusBadRequest          error = HTTPError(http.StatusBadRequest)
	ErrStatusNotFound            error = HTTPError(http.StatusNotFound)
	ErrStatusUnprocessableEntity error = HTTPError(http.StatusUnprocessableEntity)
	ErrStatusInternalServerError error = HTTPError(http.StatusInternalServerError)
)

// Реализуй меня.
type HTTPError int

func (e HTTPError) Error() string {
	return fmt.Sprintf("%d %s", e.Code(), http.StatusText(e.Code()))
}

func (e HTTPError) Code() int {
	return int(e)
}
