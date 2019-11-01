package main

import (
	"text/tabwriter"
)

var (
	_ valuer = (&lineupAmiibo{})
)

type lineupAmiibo struct {
	AmiiboName          string `json:"amiiboName"`
	AmiiboPage          string `json:"amiiboPage"`
	BoxArtURL           string `json:"boxArtUrl"`
	DetailsPath         string `json:"detailsPath"`
	DetailsURL          string `json:"detailsUrl"`
	FigureURL           string `json:"figureUrl"`
	Franchise           string `json:"franchise"`
	GameCode            string `json:"gameCode"`
	HexCode             string `json:"hexCode"`
	IsReleased          bool   `json:"isReleased"`
	OverviewDescription string `json:"overviewDescription"`
	PresentedBy         string `json:"presentedBy"`
	Price               string `json:"price"`
	ReleaseDateMask     string `json:"releaseDateMask"`
	Series              string `json:"series"`
	Slug                string `json:"slug"`
	Type                string `json:"type"`
	UnixTimestamp       int64  `json:"unixTimestamp"`
	UPC                 string `json:"upc"`
}

func (l *lineupAmiibo) Value() interface{} {
	return *l
}

func marshalLineupAmiibo(l *lineupAmiibo) (*[]byte, error) {
	return marshal(l)
}

func stringifyMarshalLineupAmiibo(l *lineupAmiibo) string {
	return stringifyMarshal(l)
}

func tableLineupAmiibo(w *tabwriter.Writer, l *lineupAmiibo) error {
	return printlnTable(w, *l)
}

func unmarshalLineupAmiibo(b *[]byte) (*lineupAmiibo, error) {
	var (
		err error
		l   lineupAmiibo
		ok  bool
	)
	err = unmarshal(b, &l)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &l, err
}
