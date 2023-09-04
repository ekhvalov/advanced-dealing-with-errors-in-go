package requests

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const maxPageSize = 100

// Реализуй нас.
var (
	errIsNotRegexp     = errors.New("exp is not regexp")
	errInvalidPage     = errors.New("invalid page")
	errInvalidPageSize = errors.New("invalid page size")
)

// Реализуй мои методы.
type ValidationErrors []error

func (ve *ValidationErrors) Error() string {
	sb := strings.Builder{}
	sb.WriteString("validation errors:\n")
	for _, err := range *ve {
		sb.WriteString(fmt.Sprintf("\t%v\n", err))
	}
	return sb.String()
}

func (ve *ValidationErrors) Is(err error) bool {
	for _, e := range *ve {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

type SearchRequest struct {
	Exp      string
	Page     int
	PageSize int
}

func (r SearchRequest) Validate() error {
	var vErrors ValidationErrors
	_, err := regexp.Compile(r.Exp)
	if err != nil {
		vErrors = append(vErrors, fmt.Errorf("%w: %v", errIsNotRegexp, err))
	}
	if r.Page <= 0 {
		vErrors = append(vErrors, fmt.Errorf("%w: %d", errInvalidPage, r.Page))
	}
	if r.PageSize <= 0 {
		vErrors = append(vErrors, fmt.Errorf("%w: %d <= 0", errInvalidPageSize, r.PageSize))
	}
	if r.PageSize > maxPageSize {
		vErrors = append(vErrors, fmt.Errorf("%w: %d > %d", errInvalidPageSize, r.PageSize, maxPageSize))
	}
	if len(vErrors) == 0 {
		return nil
	}
	return &vErrors
}
