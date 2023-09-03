package pipe

import (
	"errors"
	"fmt"
)

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

func IsPipelineError(err error, user, pipelineName string) bool {
	var pe *PipelineError
	switch {
	case errors.As(err, &pe):
		return pe.User == user && pe.Name == pipelineName
	}
	return false
}
