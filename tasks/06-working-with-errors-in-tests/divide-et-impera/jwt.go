package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

func parseHeader(data []byte) (h Header, err error) {
	defer func() {
		if err != nil {
			err = newMultiError(ErrInvalidHeaderEncoding, err)
		}
	}()
	b64Data := make([]byte, len(data))
	if err = decodeBase64(b64Data, data); err != nil {
		return h, err
	}

	return h, decodeJSON(b64Data, &h)
}

func parsePayload(data []byte) (t Token, err error) {
	defer func() {
		if err != nil {
			err = newMultiError(ErrInvalidPayloadEncoding, err)
		}
	}()
	b64Data := make([]byte, len(data))
	if err = decodeBase64(b64Data, data); err != nil {
		return t, err
	}

	return t, decodeJSON(b64Data, &t)
}

func decodeBase64(dst, src []byte) error {
	if _, err := base64.RawURLEncoding.Decode(dst, src); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidBase64, err)
	}
	return nil
}

func decodeJSON(src []byte, target interface{}) error {
	d := json.NewDecoder(bytes.NewReader(src))
	err := d.Decode(target)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidJSON, err)
	}
	return nil
}

func newMultiError(err1, err2 error) *multiError {
	return &multiError{err1: err1, err2: err2}
}

type multiError struct {
	err1 error
	err2 error
}

func (e *multiError) Error() string {
	return fmt.Sprintf("%v: %v", e.err1, e.err2)
}

func (e *multiError) Is(target error) bool {
	if errors.Is(e.err1, target) {
		return true
	}
	return errors.Is(e.err2, target)
}

func (e *multiError) As(target interface{}) bool {
	if errors.As(e.err1, target) {
		return true
	}
	return errors.As(e.err2, target)
}
