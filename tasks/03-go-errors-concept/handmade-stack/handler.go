package handmadestack

import (
	"errors"
	"fmt"
)

var (
	ErrExecSQL         = errors.New("exec sql error")
	ErrInitTransaction = errors.New("init transaction error")
)

type Entity struct {
	ID string
}

// Используются тестами.
var (
	getEntity        = func() (Entity, error) { return Entity{ID: "some-id"}, nil }
	updateEntity     = func(e Entity) error { return nil }
	runInTransaction = func(f func() error) error { return f() }
)

// Перепиши меня так, чтобы логика сохранилась,
// но путь до каждой ошибки был очевиден.
func handler() (Entity, error) {
	var e Entity

	if err := runInTransaction(func() (opErr error) {
		e, opErr = getEntity()
		if opErr != nil {
			return fmt.Errorf("get entity error: %v", opErr)
		}

		return updateEntity(e)
	}); err != nil {
		return Entity{}, fmt.Errorf("first update entity error: %v", err)
	}

	if err := runInTransaction(func() error {
		return updateEntity(e)
	}); err != nil {
		return Entity{}, fmt.Errorf("second update entity error: %v", err)
	}

	if err := runInTransaction(func() (opErr error) {
		return updateEntity(e)
	}); err != nil {
		return Entity{}, fmt.Errorf("third update entity error: %v", err)
	}

	return e, nil
}
