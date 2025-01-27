package queue

import (
	"io"
	"time"
)

const defaultPostpone = time.Second

type Job struct {
	ID int
}

type Handler struct{}

func (h *Handler) Handle(job Job) (postpone time.Duration, err error) {
	err = h.process(job)
	if err != nil {
		if shouldBeSkipped(err) {
			return 0, nil
		}
		if isTemporary(err) {
			return defaultPostpone, nil
		}
		return 0, err
	}

	return 0, nil
}

func isTemporary(err error) bool {
	if e, ok := err.(interface{ Temporary() bool }); ok {
		return e.Temporary()
	}
	return false
}

func shouldBeSkipped(err error) bool {
	if e, ok := err.(interface{ Skip() bool }); ok {
		return e.Skip()
	}
	return false
}

func (h *Handler) process(job Job) error {
	switch job.ID {
	case 1:
		return &InconsistentDataError{}
	case 2:
		return &NotReadyError{}
	case 3:
		return &NotFoundError{}
	case 4:
		return &AlreadyDoneError{}
	case 5:
		return &InvalidIDError{}
	case 6:
		return io.EOF
	}
	return nil
}
