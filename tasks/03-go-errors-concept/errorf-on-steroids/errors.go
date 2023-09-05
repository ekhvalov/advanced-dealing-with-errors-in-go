package errs

import (
	"errors"
	"fmt"
	"strings"
)

func Errorf(format string, args ...any) error {
	p := newPrinter(format, args)
	msg, errs := p.print()
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return &wrapError{msg: msg, err: errs[0]}
	}
	return &wrapErrors{msg: msg, errors: errs}
}

type wrapErrors struct {
	msg    string
	errors []error
}

func (e *wrapErrors) Error() string {
	return e.msg
}

func (e *wrapErrors) Unwrap() []error {
	return e.errors
}

func (e *wrapErrors) As(target interface{}) bool {
	for _, wErr := range e.errors {
		if errors.As(wErr, target) {
			return true
		}
	}
	return false
}

func (e *wrapErrors) Is(err error) bool {
	for _, wErr := range e.errors {
		if errors.Is(wErr, err) {
			return true
		}
	}
	return false
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

type runeHandler func(r rune)

type printer struct {
	format      string
	args        []interface{}
	argN        int
	msg         strings.Builder
	errs        []error
	runeHandler runeHandler
}

func newPrinter(format string, args []interface{}) *printer {
	p := printer{
		format: format,
		args:   args,
		msg:    strings.Builder{},
	}
	p.msg.Grow(len(format))
	p.runeHandler = p.handleRune
	return &p
}

func (p *printer) print() (string, []error) {
	for _, r := range p.format {
		p.runeHandler(r)
	}
	return p.msg.String(), p.errs
}

func (p *printer) handleRune(r rune) {
	if r == '%' {
		p.runeHandler = p.handleVerb
		return
	}
	p.msg.WriteRune(r)
}

func (p *printer) handleVerb(verb rune) {
	if p.argN >= len(p.args) {
		p.msg.WriteString("%!" + string(verb) + "(MISSING)")
		return
	}
	arg := p.args[p.argN]
	switch verb {
	case 's':
		p.msg.WriteString(fmt.Sprintf("%s", arg))
	case 'd':
		p.msg.WriteString(fmt.Sprintf("%d", arg))
	case 'q':
		p.msg.WriteString(fmt.Sprintf("%q", arg))
	case 'v':
		p.msg.WriteString(fmt.Sprintf("%v", arg))
	case 't':
		p.msg.WriteString(fmt.Sprintf("%t", arg))
	case 'w':
		switch t := arg.(type) {
		case error:
			p.errs = append(p.errs, t)
			p.msg.WriteString(t.Error())
		default:
			p.msg.WriteString(fmt.Sprintf("%v", arg))
		}
	case '%':
		p.msg.WriteString("%")
	default:
		p.msg.WriteString(string(verb))
	}
	p.argN++
	p.runeHandler = p.handleRune
}
