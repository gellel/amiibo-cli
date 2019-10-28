package main

import (
	"fmt"
	"sync"
)

type mix struct {
	CompatabilityAmiibo []*compatabilityAmiibo
	CompatabilityGame   []*compatabilityGame
	CompatabilityItem   []*compatabilityItem
	LineupAmiibo        []*lineupAmiibo
	LineupItem          []*lineupItem
}

func getMix() (*mix, error) {
	var (
		c   *compatability
		err error
		l   *lineup
		ok  bool
		wg  sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c, err = getCompatability()
		ok = (err == nil)
		if !ok {
			return
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		l, err = getLineup()
		ok = (err == nil)
		if !ok {
			return
		}
	}()
	wg.Wait()
	return newMix(c, l)
}

func newMix(c *compatability, l *lineup) (*mix, error) {
	var (
		err error
		m   *mix
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
	m = &mix{
		CompatabilityAmiibo: c.Amiibo,
		CompatabilityGame:   c.Games,
		CompatabilityItem:   c.Items,
		LineupAmiibo:        l.Amiibo,
		LineupItem:          l.Items}
	return m, err
}
