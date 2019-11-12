package main

import (
	"net/http"
)

type gameMuxAll struct {
	Games []byte
}

func (g *gameMuxAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		contentType      string = "content-type"
		contentTypeValue string = "application/json; charset=utf-8"
	)
	w.Header().Set(contentType, contentTypeValue)
	w.Write(g.Games)
	w.WriteHeader(http.StatusOK)
}

func newGameMuxAll(s []*game) (*gameMuxAll, error) {
	var (
		b   *[]byte
		err error
		g   gameMuxAll
		ok  bool
	)
	b, err = marshal(&s)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	g = gameMuxAll{
		Games: *b}
	return &g, err
}

func newGameMuxAllFromMix(m *mix) (*gameMuxAll, error) {
	var (
		err        error
		gameMap    *gameMap
		mixGameMap *mixGameMap
		ok         bool
	)
	mixGameMap, err = newMixGameMapFromMix(m)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	gameMap, err = newGameMap(mixGameMap)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return newGameMuxAll(gameMap.Values())
}
