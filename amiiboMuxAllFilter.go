package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type amiiboMuxAllFilter struct {
	Amiibo []*amiibo
}

func (a *amiiboMuxAllFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		contentType      string = "content-type"
		contentTypeValue string = "application/json; charset=utf-8"
	)
	const (
		sep string = ","
	)
	var (
		vars = mux.Vars(r)
	)
	var (
		b           *[]byte
		hasFilters  = len(vars) != 0
		released, _ = strconv.ParseBool(vars["released"])
		x           = []*amiibo{}
		rMap        = map[string]bool{}
		series      = strings.Split(vars["series"], sep)
		seriesMap   = map[string]bool{}
	)
	for _, series := range series {
		seriesMap[series] = true
	}
	for _, a := range a.Amiibo {
		if ok := seriesMap[a.Series]; ok {
			rMap[a.URI] = true
		}
		if ok := a.IsReleased == released; ok {
			rMap[a.URI] = true
		}
		if hasFilters {
			if _, ok := rMap[a.URI]; ok {
				x = append(x, a)
			}
		} else {
			x = append(x, a)
		}
	}
	fmt.Println(x)
	b, err := marshal(&x)
	if err != nil {
		panic(err)
	}
	w.Header().Set(contentType, contentTypeValue)
	w.Write(*b)
	w.WriteHeader(http.StatusOK)
}

func newAmiiboMuxAllFilter(s []*amiibo) (*amiiboMuxAllFilter, error) {
	var (
		a   amiiboMuxAllFilter
		err error
	)
	a = amiiboMuxAllFilter{
		Amiibo: s}
	return &a, err
}
