package main

import (
	"fmt"
	"sync"
)

type mixAmiiboMap map[string]*mixAmiibo

func newMixAmiiboMap(c *compatabilityAmiiboMap, l *lineupAmiiboMap, i *lineupItemMap) (*mixAmiiboMap, error) {
	var (
		err error
		m   = mixAmiiboMap{}
		ok  bool

		mu sync.Mutex
		wg sync.WaitGroup
	)
	ok = (c != nil)
	if !ok {
		return nil, fmt.Errorf("*c is nil")
	}
	ok = (l != nil)
	if !ok {
		return nil, fmt.Errorf("*l is nil")
	}
	ok = (i != nil)
	if !ok {
		return nil, fmt.Errorf("*i is nil")
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for k, v := range *c {
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &mixAmiibo{}
			}
			x := m[k]
			x.compatabilityAmiibo = v
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for k, v := range *l {
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &mixAmiibo{}
			}
			x := m[k]
			x.lineupAmiibo = v
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for k, v := range *i {
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &mixAmiibo{}
			}
			x := m[k]
			x.lineupItem = v
			mu.Unlock()
		}
	}()
	wg.Wait()
	return &m, err
}

func newMixAmiiboMapFromMix(m *mix) (*mixAmiiboMap, error) {
	var (
		c  *compatabilityAmiiboMap
		l  *lineupAmiiboMap
		i  *lineupItemMap
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
	wg.Add(1)
	go func() {
		defer wg.Done()
		i, _ = newLineupItemMap(m.LineupItem)
	}()
	wg.Wait()
	return newMixAmiiboMap(c, l, i)
}
