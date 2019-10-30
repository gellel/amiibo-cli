package main

import (
	"fmt"
	"net/http"
	"text/tabwriter"
)

type lineup struct {
	Amiibo           []*lineupAmiibo `json:"amiiboList"`
	ComponentPath    string          `json:"componentPath"`
	DateFormatString string          `json:"dataFormatString"`
	Items            []*lineupItem   `json:"items"`
}

func getLineup() (*lineup, error) {
	var (
		err error
		ok  bool
		x   *xhr
	)
	x, err = net(lineupURI)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	ok = (x.StatusCode == http.StatusBadRequest)
	if !ok {
		return nil, fmt.Errorf(x.Status)
	}
	return unmarshalLineup(x.Body)
}

func fillLineup(c *lineup) *lineup {
	return c
}

func marshalLineup(l *lineup) (*[]byte, error) {
	return marshal(l)
}

func tableLineup(w *tabwriter.Writer, l *lineup) error {
	return printlnTable(w, *l)
}

func unmarshalLineup(b *[]byte) (*lineup, error) {
	var (
		c   lineup
		err error
		ok  bool
	)
	err = unmarshal(b, &c)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &c, err
}
