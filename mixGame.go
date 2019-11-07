package main

import "fmt"

type mixGame struct {
	*compatabilityItem
	*compatabilityGame
}

func newMixGame(g *compatabilityGame, i *compatabilityItem) (*mixGame, error) {
	var (
		err error
		m   mixGame
		ok  bool
	)
	ok = (g != nil)
	if !ok {
		return nil, fmt.Errorf("*g is nil")
	}
	ok = (i != nil)
	if !ok {
		return nil, fmt.Errorf("*i is nil")
	}
	ok = (g.Path == i.Path)
	if !ok {
		return nil, fmt.Errorf("*i does not relate to *g")
	}
	m = mixGame{
		compatabilityGame: g,
		compatabilityItem: i}
	return &m, err
}
