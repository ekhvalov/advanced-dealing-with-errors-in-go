package errors

// TrimStackTrace режет все стектрейсы в цепочке ошибок err.
func TrimStackTrace(err error) error {
	if nil == err {
		return nil
	}
	return &wrapper{err: err}
}

type wrapper struct {
	err error
}

func (w *wrapper) Error() string {
	return w.err.Error()
}

func (w *wrapper) Unwrap() error {
	return w.err
}
