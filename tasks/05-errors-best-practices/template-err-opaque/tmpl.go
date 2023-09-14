package tmpl

import (
	"errors"
	"regexp"
	"text/template"
)

var funcNotDefinedRe = regexp.MustCompile(`template: \w+:\d+: function "\w+" not defined`)

func IsFunctionNotDefinedError(err error) bool {
	return err != nil && funcNotDefinedRe.MatchString(err.Error())
}

var unexportedFieldRe = regexp.MustCompile(`template:.*executing.*: \w+ is an unexported field of struct type`)

func IsExecUnexportedFieldError(err error) bool {
	var execErr template.ExecError
	return errors.As(err, &execErr) && unexportedFieldRe.MatchString(execErr.Error())
}
