package main

import (
	"fmt"
)

type mixAmiiboMap map[string]*mixAmiibo

func newMixAmiiboMap(c *compatabilityAmiiboMap, l *lineupAmiiboMap) (*mixAmiiboMap, error) {
	var (
		err error
		m   = mixAmiiboMap{}
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
	fmt.Println(len(*c), len(*l))
	for k, v := range *c {
		m[k] = &mixAmiibo{compatabilityAmiibo: v}
	}
	for k, v := range *l {
		_, exists := m[k]
		if !exists {
			m[k] = &mixAmiibo{lineupAmiibo: v}
		} else {
			m[k].lineupAmiibo = v
		}
	}
	return &m, err
}
