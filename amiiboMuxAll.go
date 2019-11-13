package main

import "net/http"

type amiiboMuxAll struct {
	Amiibo []byte
}

func (a amiiboMuxAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		contentType      string = "content-type"
		contentTypeValue string = "application/json; charset=utf-8"
	)
	w.Header().Set(contentType, contentTypeValue)
	w.Write(a.Amiibo)
	w.WriteHeader(http.StatusOK)
}

func newAmiiboMuxAll(s []*amiibo) (*amiiboMuxAll, error) {
	var (
		a   amiiboMuxAll
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshal(&s)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	a = amiiboMuxAll{
		Amiibo: *b}
	return &a, err
}
