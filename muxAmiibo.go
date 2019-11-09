package main

import "net/http"

type muxAmiiboHandler struct {
	Amiibo *amiibo
}

func (m muxAmiiboHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader()
}
