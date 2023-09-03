package pipe

import "fmt"

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

// Добавь метод Is для типа *PipelineError.
func (p *PipelineError) Is(err error) bool {
	e, ok := err.(*PipelineError)
	if ok {
		return p.User == e.User && p.Name == e.Name
	}
	return false
}
