package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type gameMuxName struct {
	Games *gameMap
}

func (m gameMuxName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		key string = "name"
	)
	var (
		b    *[]byte
		err  error
		g    *game
		name = mux.Vars(r)[key]
		ok   bool
	)
	g, ok = m.Games.Get(name)
	if !ok {
		return
	}
	b, err = marshalGame(g)
	ok = (err == nil)
	if !ok {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(*b)
}
