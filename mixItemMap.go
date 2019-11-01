package main

import "fmt"

type mixItemMap map[string]*mixItem

func newMixItemMap(c *compatabilityItemMap, l *lineupItemMap) (*mixItemMap, error) {
	var (
		err error
		m   = mixItemMap{}
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
	for k, v := range *c {
		m[k] = &mixItem{compatabilityItem: v}
	}
	for k, v := range *l {
		_, exists := m[k]
		if !exists {
			m[k] = &mixItem{lineupItem: v}
		} else {
			m[k].lineupItem = v
		}
	}
	return &m, err
}
