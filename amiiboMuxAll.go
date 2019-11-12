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
