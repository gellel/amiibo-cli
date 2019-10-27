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

func intersectionMixAmiibo(m *mix) sync.Map {
	var (
		mu sync.Map
		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, c := range m.CompatabilityAmiibo {
			syncMuKeyCount(&mu, c.URL)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, c := range m.LineupAmiibo {
			syncMuKeyCount(&mu, c.DetailsURL)
		}
	}()
	wg.Wait()
	return mu
}

func intersectionMixItems(m *mix) sync.Map {
	var (
		mu sync.Map
		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, c := range m.CompatabilityItem {
			syncMuKeyCount(&mu, c.URL)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, c := range m.LineupItem {
			syncMuKeyCount(&mu, c.URL)
		}
	}()
	wg.Wait()
	return mu
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
