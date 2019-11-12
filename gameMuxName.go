package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type gameMuxName struct {
	*byteMap
}

func (m gameMuxName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		contentType      string = "content-type"
		contentTypeValue string = "application/json; charset=utf-8"
		key              string = "name"
	)
	var (
		b    *[]byte
		name = mux.Vars(r)[key]
		ok   bool
	)
	b, ok = m.Get(name)
	if !ok {
		return
	}
	w.Header().Set(contentType, contentTypeValue)
	w.Write(*b)
	w.WriteHeader(http.StatusOK)
}

func newGameMuxName(m *gameMap) (*gameMuxName, error) {
	var (
		b   *byteMap
		err error
		g   gameMuxName
		ok  bool
	)
	b, err = newByteMap(m)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	g = gameMuxName{b}
	return &g, err
}

func newGameMuxNameFromMix(m *mix) (*gameMuxName, error) {
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
	return newGameMuxName(gameMap)
}
