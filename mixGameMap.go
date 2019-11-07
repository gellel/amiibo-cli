package main

import (
	"fmt"
	"sync"
)

type mixGameMap map[string]*mixGame

func newMixGameMap(c *compatabilityGameMap, i *compatabilityItemMap) (*mixGameMap, error) {
	var (
		err error
		m   = mixGameMap{}
		ok  bool

		mu sync.Mutex
		wg sync.WaitGroup
	)
	ok = (c != nil)
	if !ok {
		return nil, fmt.Errorf("*c is nil")
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
				m[k] = &mixGame{}
			}
			x := m[k]
			x.compatabilityGame = v
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for k, v := range *i {
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &mixGame{}
			}
			x := m[k]
			x.compatabilityItem = v
			mu.Unlock()
		}
	}()
	wg.Wait()
	return &m, err
}
