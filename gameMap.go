package main

import "fmt"

type gameMap map[string]*game

func newGameMap(m *mixGameMap) (*gameMap, error) {
	var (
		err error
		//g   *game
		ok bool

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
		newGame(v.compatabilityGame, v.compatabilityItem)
	}
	return &x, err
}
