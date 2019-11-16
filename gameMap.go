package main

import (
	"fmt"
	"sync"
)

var (
	_ hashMap = &gameMap{}
)

type gameMap map[string]*game

func (m *gameMap) Del(key string) bool {
	delete((*m), key)
	return (m.Has(key) == false)
}

func (m *gameMap) Each(fn func(string, interface{})) {
	for key, value := range *m {
		fn(key, value)
	}
}

func (m *gameMap) Get(key string) (*game, bool) {
	var (
		game *game
		ok   bool
	)
	game, ok = ((*m)[key])
	return game, ok
}

func (m *gameMap) Has(key string) bool {
	var (
		ok bool
	)
	_, ok = m.Get(key)
	return ok
}

func (m *gameMap) Keys() []string {
	var (
		keys []string
	)
	m.Each(func(key string, _ interface{}) {
		keys = append(keys, key)
	})
	return keys
}

func (m *gameMap) Len() int {
	return len(*m)
}

func (m *gameMap) Values() []*game {
	var (
		games []*game
	)
	for _, game := range *m {
		games = append(games, game)
	}
	return games
}

func newGameMap(m *mixGameMap) (*gameMap, error) {
	var (
		err error
		g   *game
		mu  sync.Mutex
		ok  bool
		wg  sync.WaitGroup
		x   = gameMap{}
	)
	ok = (m != nil)
	if !ok {
		return nil, fmt.Errorf("*m is nil")
	}
	ok = (len(*m) != 0)
	if !ok {
		return nil, fmt.Errorf("*m is empty")
	}
	for _, v := range *m {
		wg.Add(1)
		go func(v *mixGame) {
			defer wg.Done()
			g, err = newGame(v.compatabilityGame, v.compatabilityItem)
			if err != nil {
				return
			}
			if g == nil {
				return
			}
			mu.Lock()
			x[g.URI] = g
			mu.Unlock()
		}(v)
	}
	wg.Wait()
	return &x, err
}

func newGameMapFromMix(m *mix) (*gameMap, error) {
	var (
		err        error
		mixGameMap *mixGameMap
		ok         bool
	)
	mixGameMap, err = newMixGameMapFromMix(m)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return newGameMap(mixGameMap)
}
