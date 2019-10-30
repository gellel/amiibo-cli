package main

import "fmt"

type mixItem struct {
	*compatabilityItem
	*lineupItem
}

func newMixItem(c *compatabilityItem, l *lineupItem) (*mixItem, error) {
	var (
		err error
		m   mixItem
		ok  bool
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
		return nil, fmt.Errorf("*c does not relate to *l")
	}
	m = mixItem{c, l}
	return &m, err
}
