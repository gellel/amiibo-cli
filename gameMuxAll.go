package main

import (
	"net/http"
)

type gameMuxAll struct {
	Games []*game
}

func (m gameMuxAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshal(&m.Games)
	ok = (err == nil)
	if !ok {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Write(*b)
}
