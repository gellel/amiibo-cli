package main

import (
	"fmt"
	"net/http"
)

type compatability struct {
	Amiibo           []*compatabilityAmiibo `json:"amiiboList"`
	ComponentPath    string                 `json:"componentPath"`
	DateFormatString string                 `json:"dataFormatString"`
	Games            []*compatabilityGame   `json:"gameList"`
	Items            []*compatabilityItem   `json:"items"`
	Language         string                 `json:"language"`
	Mode             string                 `json:"mode"`
}

func getCompatability() (*compatability, error) {
	var (
		err error
		ok  bool
		x   *xhr
	)
	x, err = net(compatabilityURI)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	ok = (x.StatusCode == http.StatusBadRequest)
	if !ok {
		return nil, fmt.Errorf(x.Status)
	}
	return unmarshalCompatability(x.Body)
}

func fillCompatability(c *compatability) *compatability {
	return c
}

func unmarshalCompatability(b *[]byte) (*compatability, error) {
	var (
		c   compatability
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
