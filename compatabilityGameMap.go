package main

import "fmt"

type compatabilityGameMap map[string]*compatabilityGame

func newCompatabilityGameMap(c []*compatabilityGame) (*compatabilityGameMap, error) {
	var (
		err error
		m   = compatabilityGameMap{}
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
		m[v.URL] = v
	}
	return &m, err
}
