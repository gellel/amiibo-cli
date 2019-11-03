package main

import "fmt"

type mixAmiibo struct {
	*compatabilityAmiibo
	*lineupAmiibo
	*lineupItem
}

func newMixAmiibo(c *compatabilityAmiibo, l *lineupAmiibo) (*mixAmiibo, error) {
	var (
		err error
		m   mixAmiibo
		ok  bool
	)
	if !ok {
		return nil, fmt.Errorf("*c is nil")
	}
	ok = (l != nil)
	if !ok {
		return nil, fmt.Errorf("*l is nil")
	}
	ok = (c.URL == l.DetailsURL)
	if !ok {
		return nil, fmt.Errorf("*c does not relate to *l")
	}
	m = mixAmiibo{compatabilityAmiibo: c, lineupAmiibo: l}
	return &m, err
}
