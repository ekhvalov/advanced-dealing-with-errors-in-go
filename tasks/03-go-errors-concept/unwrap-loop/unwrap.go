package errs

type Unwrapper interface {
	Unwrap() error
}

func Unwrap(err error) error {
	for {
		u, ok := err.(Unwrapper)
		if !ok {
			break
		}
		err = u.Unwrap()
	}
	return err
}
