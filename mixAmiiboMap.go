package main

import (
	"fmt"
	"sync"
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

func newMixAmiiboMapFromMix(m *mix) (*mixAmiiboMap, error) {
	var (
		c  *compatabilityAmiiboMap
		l  *lineupAmiiboMap
		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c, _ = newCompatabilityAmiiboMap(m.CompatabilityAmiibo)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		l, _ = newLineupAmiiboMap(m.LineupAmiibo)
	}()
	wg.Wait()
	return newMixAmiiboMap(c, l)
}
