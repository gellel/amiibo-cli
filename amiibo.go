package main

type amiibo struct {
	BoxArtURL           string `json:"box_art_URL"`
	DetailsPath         string `json:"details_path"`
	DetailsURL          string `json:"details_URL"`
	FigureURL           string `json:"figure_URL"`
	Franchise           string `json:"franchise"`
	GameCode            string `json:"game_code"`
	HexCode             string `json:"hex_code"`
	ID                  string `json:"id"`
	Image               string `json:"image"`
	IsRelatedTo         string `json:"is_related_to"`
	IsReleased          string `json:"is_released"`
	Name                string `json:"name"`
	OverviewDescription string `json:"overview_description"`
	PageURL             string `json:"page"`
	PresentedBy         string `json:"presented_by"`
	Price               string `json:"price"`
	ReleaseDateMask     string `json:"release_date_mask"`
	Series              string `json:"series"`
	Slug                string `json:"slug"`
	TagID               string `json:"tagid"`
	Type                string `json:"type"`
	UnixTimestamp       int64  `json:"unix_timestamp"`
	UPC                 string `json:"UPC"`
	URL                 string `json:"URL"`
}

func newAmiibo(c *compatabilityAmiibo, l *lineupAmiibo) *amiibo {
	return &amiibo{
		BoxArtURL:           l.BoxArtURL,
		DetailsPath:         l.DetailsPath,
		DetailsURL:          l.DetailsURL,
		FigureURL:           l.FigureURL,
		Franchise:           l.Franchise,
		GameCode:            l.GameCode,
		HexCode:             l.HexCode,
		ID:                  c.ID,
		Image:               c.Image,
		IsRelatedTo:         c.IsRelatedTo,
		IsReleased:          c.IsReleased,
		Name:                c.Name,
		OverviewDescription: l.OverviewDescription,
		PageURL:             l.AmiiboPage,
		PresentedBy:         l.PresentedBy,
		Price:               l.Price,
		ReleaseDateMask:     c.ReleaseDateMask,
		Series:              l.Series,
		Slug:                l.Slug,
		TagID:               c.TagID,
		Type:                c.Type,
		UPC:                 l.UPC,
		URL:                 c.URL}
}
