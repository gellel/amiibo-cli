package main

import "fmt"

type compatabilityAmiiboMap map[string]*compatabilityAmiibo

func newCompatabilityAmiiboMap(c []*compatabilityAmiibo) (*compatabilityAmiiboMap, error) {
	var (
		err error
		m   = compatabilityAmiiboMap{}
		ok  bool
	)
	ok = (c != nil)
	if !ok {
		return nil, fmt.Errorf("*c is nil")
	}
	ok = (len(c) != 0)
	if !ok {
		return nil, fmt.Errorf("*c is empty")
	}
	for _, v := range c {
		var (
			s = normalizeAmiiboMapKey(v.URL)
		)
		m[s] = v
	}
	return &m, err
}
