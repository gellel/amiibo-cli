package main

import (
	"fmt"
	"html"
	"strings"
	"text/tabwriter"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

var (
	_ valuer = (&amiibo{})
)

type amiibo struct {
	BoxArtURL       *address     `json:"box_art_url"`
	Currency        string       `json:"currency"`
	Description     string       `json:"description"`
	DetailsPath     string       `json:"details_path"`
	DetailsURL      *address     `json:"details_url"`
	FigureURL       *address     `json:"figure_url"`
	Franchise       string       `json:"franchise"`
	GameCode        string       `json:"game_code"`
	HexCode         string       `json:"hex_code"`
	ID              string       `json:"id"`
	ImageURL        *address     `json:"image_url"`
	IsRelatedTo     string       `json:"is_related_to"`
	IsReleased      bool         `json:"is_released"`
	Language        language.Tag `json:"language"`
	Name            string       `json:"name"`
	PageURL         *address     `json:"page"`
	PresentedBy     string       `json:"presented_by"`
	Price           string       `json:"price"`
	ReleaseDateMask string       `json:"release_date_mask"`
	Series          string       `json:"series"`
	Slug            string       `json:"slug"`
	TagID           string       `json:"tag_id"`
	Timestamp       time.Time    `json:"timestamp"`
	Type            string       `json:"type"`
	UnixTimestamp   int64        `json:"unix_timestamp"`
	UPC             string       `json:"upc"`
	URL             *address     `json:"url"`
}

func (a *amiibo) Value() interface{} {
	return *a
}

func marshalAmiibo(a *amiibo) (*[]byte, error) {
	return marshal(a)
}

func newAmiibo(c *compatabilityAmiibo, l *lineupAmiibo) (*amiibo, error) {
	const (
		template string = "%s%s"
	)
	var (
		a              *amiibo
		boxAddress     *address
		currency       = currency.USD.String()
		detailsAddress *address
		err            error
		figureAddress  *address
		imageAddress   *address
		language       = language.AmericanEnglish
		ok             bool
		pageAddress    *address
		t              time.Time
		uAddress       *address
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
		return nil, fmt.Errorf("*c and *l do not share a common url")
	}
	ok = (c.Name == l.AmiiboName)
	if !ok {
		return nil, fmt.Errorf("*c and *l do not share a common name")
	}
	boxAddress, err = newAddress(fmt.Sprintf(template, nintendoURL, l.BoxArtURL))
	ok = (err == nil)
	if !ok {
		return nil, fmt.Errorf("cannot parse %s address: err %s", "box", err.Error())
	}
	detailsAddress, err = newAddress(fmt.Sprintf(template, nintendoURL, l.DetailsURL))
	ok = (err == nil)
	if !ok {
		return nil, fmt.Errorf("cannot parse %s address: err %s", "details", err.Error())
	}
	figureAddress, err = newAddress(fmt.Sprintf(template, nintendoURL, l.FigureURL))
	ok = (err == nil)
	if !ok {
		return nil, fmt.Errorf("cannot parse %s address: err %s", "figure", err.Error())
	}
	imageAddress, err = newAddress(fmt.Sprintf(template, nintendoURL, c.Image))
	ok = (err == nil)
	if !ok {
		return nil, fmt.Errorf("cannot parse %s address: err %s", "image", err.Error())
	}
	pageAddress, err = newAddress(fmt.Sprintf(template, nintendoURL, l.AmiiboPage))
	ok = (err == nil)
	if !ok {
		return nil, fmt.Errorf("cannot parse %s address: err %s", "page", err.Error())
	}
	uAddress, err = newAddress(fmt.Sprintf(template, nintendoURL, c.URL))
	ok = (err == nil)
	if !ok {
		return nil, fmt.Errorf("cannot parse %s address: err %s", "url", err.Error())
	}
	t = time.Unix(l.UnixTimestamp, 0).UTC()
	a = &amiibo{
		BoxArtURL:       boxAddress,
		Currency:        currency,
		Description:     stripAmiiboHTML(l.OverviewDescription),
		DetailsPath:     l.DetailsPath,
		DetailsURL:      detailsAddress,
		FigureURL:       figureAddress,
		Franchise:       l.Franchise,
		GameCode:        l.GameCode,
		HexCode:         l.HexCode,
		ID:              c.ID,
		ImageURL:        imageAddress,
		IsRelatedTo:     c.IsRelatedTo,
		IsReleased:      l.IsReleased,
		Language:        language,
		Name:            stripAmiiboName(c.Name),
		PageURL:         pageAddress,
		PresentedBy:     l.PresentedBy,
		Price:           l.Price,
		ReleaseDateMask: c.ReleaseDateMask,
		Series:          l.Series,
		Slug:            l.Slug,
		TagID:           c.TagID,
		Timestamp:       t,
		Type:            c.Type,
		UnixTimestamp:   l.UnixTimestamp,
		UPC:             l.UPC,
		URL:             uAddress}
	return a, err
}

func stripAmiiboHTML(s string) string {
	s = reStripSpaces.ReplaceAllString(reStripHTML.ReplaceAllString(s, " "), " ")
	s = html.UnescapeString(strings.TrimSpace(s))
	return s
}

func stripAmiiboName(s string) string {
	return (reStripName.ReplaceAllString(s, ""))
}

func stringifyMarshalAmiibo(a *amiibo) string {
	return stringifyMarshal(a)
}

func tableAmiibo(w *tabwriter.Writer, a *amiibo) error {
	return printlnTable(w, *a)
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
