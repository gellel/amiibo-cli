package main

import (
	"fmt"
	"path/filepath"
	"text/tabwriter"
)

var (
	_ valuer = (&compatabilityGame{})
)

type compatabilityGame struct {
	Image           string `json:"image"`
	ID              string `json:"id"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	ReleaseDateMask string `json:"releaseDateMask"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}

func (c *compatabilityGame) Key() string {
	var (
		s = fmt.Sprintf("%s/", c.URL)
	)
	s = filepath.Dir(s)
	s = filepath.Base(s)
	return s
}

func (c *compatabilityGame) Value() interface{} {
	return *c
}

func marshalCompatabilityGame(c *compatabilityGame) (*[]byte, error) {
	return marshal(c)
}

func tableCompatabilityGame(w *tabwriter.Writer, c *compatabilityGame) error {
	return printlnTable(w, *c)
}

func unmarshalCompatabilityGame(b *[]byte) (*compatabilityGame, error) {
	var (
		c   compatabilityGame
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

func writeCompatabilityGame(path, folder string, c *compatabilityGame) error {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalCompatabilityGame(c)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, c.Name, b)
}
