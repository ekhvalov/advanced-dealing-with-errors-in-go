package jwt

import "errors"

var (
	ErrEmptyJWT               = errors.New("empty jwt data")
	ErrInvalidTokenFormat     = errors.New("invalid token format: 'header.payload.signature' was expected")
	ErrInvalidHeaderEncoding  = errors.New("invalid header encoding")
	ErrUnsupportedTokenType   = errors.New("unsupported token type")
	ErrUnsupportedSigningAlgo = errors.New("unsupported the signing algo")
	ErrInvalidSignature       = errors.New("invalid signature")
	ErrInvalidPayloadEncoding = errors.New("invalid payload encoding")
	ErrExpiredToken           = errors.New("token was expired")
)

func newErrorWithEmail(err error, email string) *errWithEmail {
	return &errWithEmail{err: err, email: email}
}

type errWithEmail struct {
	err   error
	email string
}

func (e *errWithEmail) Error() string {
	return e.err.Error()
}

func (e *errWithEmail) Unwrap() error {
	return e.err
}

func (e *errWithEmail) Email() string {
	return e.email
}
