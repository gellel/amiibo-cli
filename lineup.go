package main

import (
	"fmt"
	"net/http"
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

func marhsalLineup(l *lineup) (*[]byte, error) {
	return marshalB(l)
}

func unmarshalLineup(b *[]byte) (*lineup, error) {
	var (
		c   lineup
		err error
		ok  bool
	)
	err = unmarshalB(b, &c)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &c, err
}
