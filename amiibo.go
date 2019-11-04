package main

import (
	"fmt"
	"html"
	"strconv"
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
	LastModified    int64        `json:"last_modified"`
	Name            string       `json:"name"`
	Overview        string       `json:"overview"`
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

func newAmiibo(c *compatabilityAmiibo, l *lineupAmiibo, i *lineupItem) (*amiibo, error) {
	var (
		ok bool
	)
	ok = (c != nil) || (l != nil) || (i != nil)
	if !ok {
		return nil, fmt.Errorf("*c, *l and *i are nil")
	}
	var (
		a            *amiibo
		currency     = currency.USD.String()
		description  string
		franchise    string
		game         string
		hex          string
		ID           string
		isRelatedTo  string
		isReleased   bool
		language     = language.AmericanEnglish
		lastModified int64
		name         string
		overview     string
		presentedBy  string
		price        string
		tagID        string
	)

	if c != nil {
		ID = c.ID
		//c.Image
		isRelatedTo = c.IsRelatedTo
		isReleased, _ = strconv.ParseBool(c.IsReleased)
		name = regexpName.ReplaceAllString(c.Name, "")
		tagID = c.TagID
		//c.Type
		//c.URL
	}
	if l != nil {
		//l.AmiiboPage
		//l.BoxArtURL
		//l.DetailsPath
		//l.DetailsURL
		//l.FigureURL
		franchise = l.Franchise
		game = l.GameCode
		hex = l.HexCode
		isReleased = l.IsReleased
		name = regexpName.ReplaceAllString(l.AmiiboName, "")
		overview = l.OverviewDescription
		presentedBy = stripAmiiboPresentedBy(l.PresentedBy)
		price = l.Price
		//l.ReleaseDateMask
		//l.Series
		//l.Slug
		//l.Type
		//l.UPC
		//l.UnixTimestamp
	}
	if i != nil {
		description = i.Description
		lastModified = i.LastModified
		//i.Path
		name = regexpName.ReplaceAllString(i.Title, "")
		//i.URL
	}
	a = &amiibo{
		Currency:     currency,
		Description:  description,
		Franchise:    franchise,
		GameCode:     game,
		HexCode:      hex,
		ID:           ID,
		IsRelatedTo:  isRelatedTo,
		IsReleased:   isReleased,
		Language:     language,
		LastModified: lastModified,
		Name:         name,
		Overview:     overview,
		PresentedBy:  presentedBy,
		Price:        price,
		TagID:        tagID}

	return a, nil
}

func stripAmiiboHTML(s string) string {
	s = regexpSpaces.ReplaceAllString(regexpHTML.ReplaceAllString(s, " "), " ")
	s = html.UnescapeString(strings.TrimSpace(s))
	return s
}

func stripAmiiboName(s string) string {
	return (regexpName.ReplaceAllString(s, ""))
}

func stripAmiiboPresentedBy(s string) string {
	return strings.TrimPrefix(s, "noa:publisher/")
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
