package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type amiiboMuxName struct {
	*byteMap
}

func (a amiiboMuxName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		contentType      string = "content-type"
		contentTypeValue string = "application/json; charset=utf-8"
		key              string = "name"
	)
	var (
		b          *[]byte
		ok         bool
		statusCode = http.StatusNotFound

		vars = mux.Vars(r)
	)
	b, ok = a.Get(vars[key])
	if ok {
		statusCode = http.StatusOK
	}
	w.Header().Set(contentType, contentTypeValue)
	w.Write(*b)
	w.WriteHeader(statusCode)
}

func newAmiiboMuxName(m *amiiboMap) (*amiiboMuxName, error) {
	var (
		a   amiiboMuxName
		b   *byteMap
		err error
		ok  bool
	)
	b, err = newByteMap(m)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	a = amiiboMuxName{b}
	return &a, err
}

func newAmiiboMuxNameFromMix(m *mix) (*amiiboMuxName, error) {
	var (
		err          error
		amiiboMap    *amiiboMap
		mixAmiiboMap *mixAmiiboMap
		ok           bool
	)
	mixAmiiboMap, err = newMixAmiiboMapFromMix(m)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	amiiboMap, err = newAmiiboMap(mixAmiiboMap)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return newAmiiboMuxName(amiiboMap)
}
