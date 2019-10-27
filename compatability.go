package main

import (
	"fmt"
	"net/http"
	"text/tabwriter"
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

func marshalCompatability(c *compatability) (*[]byte, error) {
	return marshalB(c)
}

func tableCompatability(w *tabwriter.Writer, c *compatability) error {
	var (
		x = &[]string{
			fmt.Sprintf("amiibo (n)\t%d", len(c.Amiibo)),
			fmt.Sprintf("component path\t%s", c.ComponentPath),
			fmt.Sprintf("date format string\t%s", c.DateFormatString),
			fmt.Sprintf("games (n)\t%d", len(c.Games)),
			fmt.Sprintf("items (n)\t%d", len(c.Items)),
			fmt.Sprintf("language\t%s", c.Language),
			fmt.Sprintf("mode\t%s", c.Mode)}
	)
	return table(w, x)
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
