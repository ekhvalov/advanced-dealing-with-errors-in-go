package errs

// Extract достаёт из цепочки err набор sentinel-ошибок,
// игнорируя "оборачивающие" их ошибки.
func Extract(err error) []error {
	errs := make([]error, 0)
	switch x := err.(type) {
	case interface{ Unwrap() error }:
		e := Extract(x.Unwrap())
		if e != nil {
			errs = append(errs, e...)
		}
	case interface{ Unwrap() []error }:
		for _, e := range x.Unwrap() {
			er := Extract(e)
			if er != nil {
				errs = append(errs, er...)
			}
		}
	case nil:
	default:
		errs = append(errs, err)
	}
	return errs
}
