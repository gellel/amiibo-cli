package main

import (
	"text/tabwriter"
)

var (
	_ valuer = (&compatabilityAmiibo{})
)

type compatabilityAmiibo struct {
	ID              string `json:"id"`
	Image           string `json:"image"`
	IsRelatedTo     string `json:"isRelatedTo"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	ReleaseDateMask string `json:"releaseDateMask"`
	TagID           string `json:"tagid"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}

func (c *compatabilityAmiibo) Value() interface{} {
	return *c
}

func marshalCompatabilityAmiibo(c *compatabilityAmiibo) (*[]byte, error) {
	return marshal(c)
}

func stringifyMarshalCompatabilityAmiibo(c *compatabilityAmiibo) string {
	return stringifyMarshal(c)
}

func tableCompatabilityAmiibo(w *tabwriter.Writer, c *compatabilityAmiibo) error {
	return printlnTable(w, *c)
}

func unmarshalCompatabilityAmiibo(b *[]byte) (*compatabilityAmiibo, error) {
	var (
		c   compatabilityAmiibo
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
