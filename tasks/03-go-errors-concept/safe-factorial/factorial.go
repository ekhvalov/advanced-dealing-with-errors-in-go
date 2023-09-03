package factorial

import (
	"errors"
	"fmt"
)

const maxDepth = 256

// Реализуй нас.
var (
	ErrNegativeN = errors.New("negative number")
	ErrTooDeep   = errors.New("call stack is too deep")
)

// Calculate рекурсивно считает факториал входного числа n.
// Если число меньше нуля, то возвращается ошибка ErrNegativeN.
// Если для вычисления факториала потребуется больше maxDepth фреймов, то Calculate вернёт ErrTooDeep.
func Calculate(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("%w: %d", ErrNegativeN, n)
	}
	if n > maxDepth {
		return 0, fmt.Errorf("%w: %d", ErrTooDeep, maxDepth)
	}
	if n < 2 {
		return 1, nil
	}
	m, err := Calculate(n - 1)
	if err != nil {
		return 0, err
	}
	return n * m, nil
}
