package main

import (
	"fmt"
	"text/tabwriter"
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

func tableLineupAmiibo(w *tabwriter.Writer, l *lineupAmiibo) error {
	var (
		x = &[]string{
			fmt.Sprintf("amiibo name\t%s", l.AmiiboName),
			fmt.Sprintf("amiibo page\t%s", l.AmiiboPage),
			fmt.Sprintf("box art url\t%s", l.BoxArtURL),
			fmt.Sprintf("details path\t%s", l.DetailsPath),
			fmt.Sprintf("details url\t%s", l.DetailsURL),
			fmt.Sprintf("figure url\t%s", l.FigureURL),
			fmt.Sprintf("franchise\t%s", l.Franchise),
			fmt.Sprintf("game code\t%s", l.GameCode),
			fmt.Sprintf("hex code\t%s", l.HexCode),
			fmt.Sprintf("is released\t%t", l.IsReleased),
			fmt.Sprintf("overview description\t%s", l.OverviewDescription),
			fmt.Sprintf("presented by\t%s", l.PresentedBy),
			fmt.Sprintf("price\t%s", l.Price),
			fmt.Sprintf("release date mask\t%s", l.ReleaseDateMask),
			fmt.Sprintf("series\t%s", l.Series),
			fmt.Sprintf("slug\t%s", l.Slug),
			fmt.Sprintf("type\t%s", l.Type),
			fmt.Sprintf("unix timestamp\t%d", l.UnixTimestamp),
			fmt.Sprintf("upc\t%s", l.UPC)}
	)
	return printlnTable(w, x)
}
