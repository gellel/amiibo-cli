package main

import (
	"text/tabwriter"
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
