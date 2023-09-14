package tmpl

import (
	"errors"
	"fmt"
	"html/template"
	"io"
)

var (
	errParseTemplate   = errors.New("parse error")
	errExecuteTemplate = errors.New("execute error")
)

func ParseAndExecuteTemplate(wr io.Writer, name, text string, data any) error {
	t, err := template.New(name).Parse(text)
	if err != nil {
		return fmt.Errorf("%w: %v", errParseTemplate, err.Error())
	}
	err = t.Execute(wr, data)
	if err != nil {
		return fmt.Errorf("%w: %v", errExecuteTemplate, err.Error())
	}
	return nil
}
