package miniword

import (
	"errors"
	"fmt"
	"io"
)

const maxPages = 3

var (
	errInvalidPageNumber = errors.New("invalid page number")
	errNoMorePages       = errors.New("no more pages")
	errEmptyText         = errors.New("empty text")
)

type Document struct {
	pageNum int
	pages   []string
	err     error
}

func NewDocument() *Document {
	return &Document{
		pageNum: 0,
		pages:   make([]string, 1, maxPages),
	}
}

func (d *Document) AddPage() {
	if d.err != nil {
		return
	}
	if len(d.pages) >= maxPages {
		d.err = errNoMorePages
		return
	}
	d.pages = append(d.pages, "")
}

func (d *Document) SetActivePage(number int) {
	if d.err != nil {
		return
	}
	if number < 1 || number > maxPages {
		d.err = errInvalidPageNumber
		return
	}
	d.pageNum = number - 1
}

func (d *Document) WriteText(s string) {
	if d.err != nil {
		return
	}
	if len(s) == 0 {
		d.err = errEmptyText
	}
	d.pages[d.pageNum] = s
}

func (d *Document) WriteTo(w io.Writer) (int64, error) {
	if d.err != nil {
		return 0, d.err
	}
	n := 0
	for num, page := range d.pages {
		if len(page) > 0 {
			nn, err := fmt.Fprintf(w, "--- Page %d ---\n%s\n", num+1, page)
			n += nn
			if err != nil {
				return int64(n), err
			}
		}

	}
	return int64(n), nil
}
