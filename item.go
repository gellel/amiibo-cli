package main

import "fmt"

var (
	_ valuer = (&item{})
)

type item struct{}

func (i *item) Value() interface{} {
	return *i
}

func newItem(c *compatabilityItem, l *lineupItem) (*item, error) {
	var (
		ok bool
	)
	ok = (c != nil)
	if !ok {
		return nil, fmt.Errorf("*c is nil")
	}
	ok = (l != nil)
	if !ok {
		return nil, fmt.Errorf("*l is nil")
	}
	ok = (c.URL == l.URL)
	if !ok {
		return nil, fmt.Errorf("*c and *l do not share a common url")
	}
	ok = (c.Title == l.Title)
	if !ok {
		return nil, fmt.Errorf("*c and *l do not share a common title")
	}
	return nil, nil
}
