package errs

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
)

func BenchmarkPretty(b *testing.B) {
	firstErr := sql.ErrNoRows
	err := fmt.Errorf("cannot get user schedule: %w",
		fmt.Errorf("cannot build data for event: %w",
			fmt.Errorf("cannot get image: %w",
				fmt.Errorf("cannot get image: %w",
					fmt.Errorf("cannot get image: %w",
						fmt.Errorf("cannot get image for event: %w", firstErr))))))
	for i := 0; i < b.N; i++ {
		pErr := Pretty(err)
		errors.Is(pErr, firstErr)
	}
}
