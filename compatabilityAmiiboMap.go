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
		return nil, fmt.Errorf("*c is ni")
	}
	ok = (len(c) != 0)
	if !ok {
		return nil, fmt.Errorf("*c is empty")
	}
	for _, v := range c {
		m[v.URL] = v
	}
	return &m, err
}
