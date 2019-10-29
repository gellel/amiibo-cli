package main

import (
	"fmt"
	"time"
)

type amiibo struct {
	BoxArtURL           *addr     `json:"box_art_URL"`
	DetailsPath         string    `json:"details_path"`
	DetailsURL          *addr     `json:"details_URL"`
	FigureURL           *addr     `json:"figure_URL"`
	Franchise           string    `json:"franchise"`
	GameCode            string    `json:"game_code"`
	HexCode             string    `json:"hex_code"`
	ID                  string    `json:"id"`
	Image               string    `json:"image"`
	IsRelatedTo         string    `json:"is_related_to"`
	IsReleased          string    `json:"is_released"`
	Name                string    `json:"name"`
	OverviewDescription string    `json:"overview_description"`
	PageURL             *addr     `json:"page"`
	PresentedBy         string    `json:"presented_by"`
	Price               string    `json:"price"`
	ReleaseDateMask     string    `json:"release_date_mask"`
	Series              string    `json:"series"`
	Slug                string    `json:"slug"`
	TagID               string    `json:"tag_id"`
	Timestamp           time.Time `json:"timestamp"`
	Type                string    `json:"type"`
	UnixTimestamp       int64     `json:"unix_timestamp"`
	UPC                 string    `json:"UPC"`
	URL                 *addr     `json:"URL"`
}

func marshalAmiibo(a *amiibo) (*[]byte, error) {
	return marshal(a)
}

func newAmiibo(c *compatabilityAmiibo, l *lineupAmiibo) (*amiibo, error) {
	var (
		a           *amiibo
		boxAddr     *addr
		detailsAddr *addr
		err         error
		figureAddr  *addr
		ok          bool
		pageAddr    *addr
		t           time.Time
		uAddr       *addr
	)
	ok = (c != nil)
	if !ok {
		return nil, fmt.Errorf("*c is nil")
	}
	ok = (l != nil)
	if !ok {
		return nil, fmt.Errorf("*l is nil")
	}
	ok = (c.URL == l.DetailsURL)
	if !ok {
		return nil, fmt.Errorf("*c and *l do not share a common URL")
	}
	ok = (c.Name == l.AmiiboName)
	if !ok {
		return nil, fmt.Errorf("*c and *l do not share a common name")
	}
	boxAddr, err = newAddr(fmt.Sprintf("%s%s", nintendoURL, l.BoxArtURL))
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	detailsAddr, err = newAddr(fmt.Sprintf("%s%s", nintendoURL, l.DetailsURL))
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	figureAddr, err = newAddr(fmt.Sprintf("%s%s", nintendoURL, l.FigureURL))
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	pageAddr, err = newAddr(fmt.Sprintf("%s%s", nintendoURL, l.AmiiboPage))
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	uAddr, err = newAddr(fmt.Sprintf("%s%s", nintendoURL, c.URL))
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	t = time.Unix(0, l.UnixTimestamp)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	a = &amiibo{
		BoxArtURL:           boxAddr,
		DetailsPath:         l.DetailsPath,
		DetailsURL:          detailsAddr,
		FigureURL:           figureAddr,
		Franchise:           l.Franchise,
		GameCode:            l.GameCode,
		HexCode:             l.HexCode,
		ID:                  c.ID,
		Image:               c.Image,
		IsRelatedTo:         c.IsRelatedTo,
		IsReleased:          c.IsReleased,
		Name:                c.Name,
		OverviewDescription: l.OverviewDescription,
		PageURL:             pageAddr,
		PresentedBy:         l.PresentedBy,
		Price:               l.Price,
		ReleaseDateMask:     c.ReleaseDateMask,
		Series:              l.Series,
		Slug:                l.Slug,
		TagID:               c.TagID,
		Timestamp:           t,
		Type:                c.Type,
		UnixTimestamp:       l.UnixTimestamp,
		UPC:                 l.UPC,
		URL:                 uAddr}
	return a, err
}

func unmarshalAmiibo(b *[]byte) (*amiibo, error) {
	var (
		a   amiibo
		err error
		ok  bool
	)
	err = unmarshal(b, &a)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &a, err
}

func writeAmiibo(path, folder string, a *amiibo) error {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalAmiibo(a)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, a.Name, b)
}
