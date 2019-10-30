package main

import "text/tabwriter"

type lineupItem struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
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
