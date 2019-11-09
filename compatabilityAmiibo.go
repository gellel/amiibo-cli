package main

import (
	"text/tabwriter"
)

var (
	_ keyspace = (&compatabilityAmiibo{})
	_ valuer   = (&compatabilityAmiibo{})
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

func (c *compatabilityAmiibo) Key() string {
	return normalizeAmiiboMapKey(c.URL)
}

func (c *compatabilityAmiibo) Value() interface{} {
	return *c
}

func marshalCompatabilityAmiibo(c *compatabilityAmiibo) (*[]byte, error) {
	return marshal(c)
}

func readCompatabilityAmiibo(fullpath string) (*compatabilityAmiibo, error) {
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
	return unmarshalCompatabilityAmiibo(b)
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

func writeCompatabilityAmiibo(path, folder string, c *compatabilityAmiibo) error {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalCompatabilityAmiibo(c)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, c.Name, b)
}
