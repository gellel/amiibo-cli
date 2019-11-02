package main

import "text/tabwriter"

var (
	_ valuer = (&lineupItem{})
)

type lineupItem struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

func (l *lineupItem) Value() interface{} {
	return *l
}

func marshalLineupItem(l *lineupItem) (*[]byte, error) {
	return marshal(l)
}

func tableLineupItem(w *tabwriter.Writer, l *lineupItem) error {
	return printlnTable(w, *l)
}

func unmarshalLineupItem(b *[]byte) (*lineupItem, error) {
	var (
		err error
		l   lineupItem
		ok  bool
	)
	err = unmarshal(b, &l)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &l, err
}

func writeLineupItem(path, folder string, l *lineupItem) error {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalLineupItem(l)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, l.Title, b)
}
