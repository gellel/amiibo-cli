package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

var (
	_ valuer = (&compatabilityItem{})
)

type compatabilityItem struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

func (c *compatabilityItem) Key() string {
	var (
		s = fmt.Sprintf("%s/", strings.TrimSuffix(c.URL, ".html"))
	)
	s = filepath.Dir(s)
	s = filepath.Base(s)
	return s
}

func (c *compatabilityItem) Value() interface{} {
	return *c
}

func marshalCompatabilityItem(c *compatabilityItem) (*[]byte, error) {
	return marshal(c)
}

func tableCompatabilityItem(w *tabwriter.Writer, c *compatabilityItem) error {
	return printlnTable(w, *c)
}

func unmarshalCompatabilityItem(b *[]byte) (*compatabilityItem, error) {
	var (
		c   compatabilityItem
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

func writeCompatabilityItem(path, folder string, c *compatabilityItem) error {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalCompatabilityItem(c)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, c.Title, b)
}
