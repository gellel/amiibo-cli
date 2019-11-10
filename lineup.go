package main

import (
	"fmt"
	"net/http"
	"text/tabwriter"
)

var (
	_ valuer = (&lineup{})
)

type lineup struct {
	Amiibo           []*lineupAmiibo `json:"amiiboList"`
	ComponentPath    string          `json:"componentPath"`
	DateFormatString string          `json:"dataFormatString"`
	Items            []*lineupItem   `json:"items"`
}

func (l *lineup) Value() interface{} {
	return *l
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

func getAndWriteLineup(path, folder string) error {
	var (
		err error
		l   *lineup
		ok  bool
	)
	l, err = getLineup()
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeLineup(path, folder, l)
}

func fillLineup(c *lineup) *lineup {
	return c
}

func marshalLineup(l *lineup) (*[]byte, error) {
	return marshal(l)
}

func readLineup(fullpath string) (*lineup, error) {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = readFile(fullpath)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return unmarshalLineup(b)
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

func writeLineup(path, folder string, l *lineup) error {
	const (
		name string = "lineup"
	)
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalLineup(l)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, name, b)
}
