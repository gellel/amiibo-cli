package main

import "fmt"

type compatabilityItemMap map[string]*compatabilityItem

func newCompatabilityItemMap(c []*compatabilityItem) (*compatabilityItemMap, error) {
	var (
		err error
		m   = compatabilityItemMap{}
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
