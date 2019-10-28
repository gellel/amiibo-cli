package main

import "fmt"

type lineupAmiiboMap map[string]*lineupAmiibo

func newLineupAmiiboMap(l []*lineupAmiibo) (*lineupAmiiboMap, error) {
	var (
		err error
		m   = lineupAmiiboMap{}
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
		m[v.DetailsURL] = v
	}
	return &m, err
}
