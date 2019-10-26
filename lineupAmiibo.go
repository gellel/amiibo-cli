package main

type lineupAmiibo struct {
	AmiiboName          string  `json:"amiiboName"`
	AmiiboPage          string  `json:"amiiboPage"`
	BoxArtURL           string  `json:"boxArtUrl"`
	DetailsPath         string  `json:"detailsPath"`
	DetailsURL          string  `json:"detailsUrl"`
	GameCode            string  `json:"gameCode"`
	FigureURL           string  `json:"figureUrl"`
	Franchise           string  `json:"franchise"`
	HexCode             string  `json:"hexCode"`
	IsReleased          bool    `json:"isReleased"`
	OverviewDescription string  `json:"overviewDescription"`
	PresentedBy         string  `json:"presentedBy"`
	Price               float64 `json:"price"`
	ReleaseDateMask     string  `json:"releaseDateMask"`
	Series              string  `json:"series"`
	Slug                string  `json:"slug"`
	Type                string  `json:"type"`
	UnixTimestamp       int64   `json:"unixTimestamp"`
	UPC                 string  `json:"upc"`
}
