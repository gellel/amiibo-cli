package main

import "fmt"

type gameMap map[string]*game

func (m *gameMap) Get(key string) (*game, bool) {
	var (
		game *game
		ok   bool
	)
	game, ok = ((*m)[key])
	return game, ok
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
		ok  bool

		x = gameMap{}
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
		g, err = newGame(v.compatabilityGame, v.compatabilityItem)
		if err != nil {
			continue
		}
		if g == nil {
			continue
		}
		x[g.URI] = g
	}
	return &x, err
}
