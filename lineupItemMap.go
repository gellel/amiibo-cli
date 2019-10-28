package main

import "fmt"

type lineupItemMap map[string]*lineupItem

func newLineupItem(l []*lineupItem) (*lineupItemMap, error) {
	var (
		err error
		m   = lineupItemMap{}
		ok  bool
	)
	ok = (l != nil)
	if !ok {
		return nil, fmt.Errorf("*l is ni")
	}
	ok = (len(l) != 0)
	if !ok {
		return nil, fmt.Errorf("*l is empty")
	}
	for _, v := range l {
		m[v.URL] = v
	}
	return &m, err
}
