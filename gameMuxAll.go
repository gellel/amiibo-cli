package main

import (
	"net/http"
)

type gameMuxAll struct {
	Games []*game
}

func (m gameMuxAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		contentType      string = "content-type"
		contentTypeValue string = "application/json; charset=utf-8"
	)
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshal(&m.Games)
	ok = (err == nil)
	switch ok {
	case true:
		w.WriteHeader(http.StatusOK)
		w.Header().Set(contentType, contentTypeValue)
		w.Write(*b)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
