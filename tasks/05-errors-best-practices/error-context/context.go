package errctx

import "errors"

type Fields map[string]any

func AppendTo(err error, fields Fields) error {
	if nil == err {
		return nil
	}
	fe := &fieldsError{err: err, fields: make(Fields, len(fields))}
	for k, v := range fields {
		fe.fields[k] = v
	}
	return fe
}

func From(err error) Fields {
	fields := make(Fields)
	for err != nil {
		if errFields, ok := err.(*fieldsError); ok {
			for k, v := range errFields.fields {
				fields[k] = v
			}
		}
		err = errors.Unwrap(err)
	}
	return fields
}

type fieldsError struct {
	err    error
	fields Fields
}

func (e *fieldsError) Error() string {
	return e.err.Error()
}

func (e *fieldsError) Unwrap() error {
	return e.err
}
