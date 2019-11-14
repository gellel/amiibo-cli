package main

import "net/http"

type amiiboMuxAllFilter struct {
	Amiibo []*amiibo
}

func (a *amiiboMuxAllFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
