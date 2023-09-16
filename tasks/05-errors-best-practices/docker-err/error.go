package docker

import (
	"regexp"
	"strings"
)

type errorCode int

const (
	errorCodeUnknown errorCode = iota
	errorCodePullAccessDenied
	errorCodeNoSuchContainer
	errorCodeContainerNotRunning
)

const (
	pullAccessDeniedMsg = "pull access denied"
	noSuchContainerMsg  = "No such container"
)

var containerNotRunningRe = regexp.MustCompile(`Container \w+ is not running`)

var _ interface {
	error
	IsPullAccessDeniedError() bool
	IsNoSuchContainerError() bool
	IsContainerNotRunningError() bool
} = (*Error)(nil)

type Error struct {
	err  error
	code errorCode
}

func newDockerError(err error) *Error {
	var errCode errorCode
	switch {
	case strings.Contains(err.Error(), pullAccessDeniedMsg):
		errCode = errorCodePullAccessDenied
	case strings.Contains(err.Error(), noSuchContainerMsg):
		errCode = errorCodeNoSuchContainer
	case containerNotRunningRe.MatchString(err.Error()):
		errCode = errorCodeContainerNotRunning
	default:
		errCode = errorCodeUnknown
	}
	return &Error{err: err, code: errCode}
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) IsPullAccessDeniedError() bool {
	return e.code == errorCodePullAccessDenied
}

func (e *Error) IsNoSuchContainerError() bool {
	return e.code == errorCodeNoSuchContainer
}

func (e *Error) IsContainerNotRunningError() bool {
	return e.code == errorCodeContainerNotRunning
}

func (e *Error) Error() string {
	return e.err.Error()
}
