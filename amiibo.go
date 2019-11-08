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
	BoxImage        *image       `json:"box_image"`
	Complete        bool         `json:"complete"`
	Currency        string       `json:"currency"`
	Description     string       `json:"description"`
	DetailsPath     string       `json:"details_path"`
	DetailsURL      *address     `json:"details_url"`
	FigureURL       *address     `json:"figure_url"`
	Franchise       string       `json:"franchise"`
	GameCode        string       `json:"game_code"`
	HexCode         string       `json:"hex_code"`
	ID              string       `json:"id"`
	Image           *image       `json:"image"`
	IsRelatedTo     string       `json:"is_related_to"`
	IsReleased      bool         `json:"is_released"`
	Language        language.Tag `json:"language"`
	LastModified    int64        `json:"last_modified"`
	Name            string       `json:"name"`
	Overview        string       `json:"overview"`
	PageURL         *address     `json:"page"`
	Path            string       `json:"path"`
	PresentedBy     string       `json:"presented_by"`
	Price           string       `json:"price"`
	ReleaseDateMask string       `json:"release_date_mask"`
	Series          string       `json:"series"`
	Slug            string       `json:"slug"`
	TagID           string       `json:"tag_id"`
	Timestamp       time.Time    `json:"timestamp"`
	Type            string       `json:"type"`
	TypeAlias       string       `json:"type_alias"`
	Unix            int64        `json:"unix"`
	UPC             string       `json:"upc"`
	URI             string       `json:"uri"`
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
	const (
		template string = "%s%s"
	)
	var (
		a               *amiibo
		boxImage        *image
		complete        bool
		currency        = currency.USD.String()
		description     string
		detailsPath     string
		detailsURL      *address
		figureURL       *address
		franchise       string
		game            string
		hex             string
		ID              string
		image           *image
		isRelatedTo     string
		isReleased      bool
		language        = language.AmericanEnglish
		lastModified    int64
		name            string
		overview        string
		pageURL         *address
		path            string
		presentedBy     string
		price           string
		releaseDateMask string
		series          string
		slug            string
		tagID           string
		timestamp       time.Time
		typeAlias       string
		typeOf          string
		unix            int64
		UPC             string
		URI             string
		URL             *address
	)
	complete = (c != nil) && (l != nil) && (i != nil)
	if c != nil {
		ID = c.ID
		image, _ = newImage(fmt.Sprintf(template, nintendoURL, c.Image))
		isRelatedTo = c.IsRelatedTo
		isReleased, _ = strconv.ParseBool(c.IsReleased)
		name = stripAmiiboName(c.Name)
		tagID = c.TagID
		typeOf = c.Type
		URL, _ = newAddress(fmt.Sprintf(template, nintendoURL, c.URL))
	}
	if l != nil {
		boxImage, _ = newImage(fmt.Sprintf(template, nintendoURL, l.BoxArtURL))
		detailsPath = l.DetailsPath
		detailsURL, _ = newAddress(fmt.Sprintf(template, nintendoURL, l.DetailsURL))
		figureURL, _ = newAddress(fmt.Sprintf(template, nintendoURL, l.FigureURL))
		franchise = l.Franchise
		game = l.GameCode
		hex = l.HexCode
		isReleased = l.IsReleased
		ok = (len(name) != 0)
		if !ok {
			name = stripAmiiboName(l.AmiiboName)
		}
		overview = stripAmiiboHTML(l.OverviewDescription)
		pageURL, _ = newAddress(fmt.Sprintf(template, nintendoURL, l.AmiiboPage))
		presentedBy = stripAmiiboPresentedBy(l.PresentedBy)
		price = l.Price
		releaseDateMask = l.ReleaseDateMask
		series = l.Series
		slug = l.Slug
		timestamp = time.Unix(l.UnixTimestamp, 0)
		typeAlias = strings.ToLower(l.Type)
		UPC = l.UPC
		unix = l.UnixTimestamp
	}
	if i != nil {
		description = i.Description
		lastModified = i.LastModified
		path = i.Path
		ok = (len(name) != 0)
		if !ok {
			name = stripAmiiboName(i.Title)
		}
		ok = (URL != nil)
		if !ok {
			URL, _ = newAddress(fmt.Sprintf(template, nintendoURL, i.URL))
		}
	}
	URI = normalizeURI(name)
	a = &amiibo{
		BoxImage:        boxImage,
		Complete:        complete,
		Currency:        currency,
		Description:     description,
		DetailsPath:     detailsPath,
		DetailsURL:      detailsURL,
		FigureURL:       figureURL,
		Franchise:       franchise,
		GameCode:        game,
		HexCode:         hex,
		ID:              ID,
		Image:           image,
		IsRelatedTo:     isRelatedTo,
		IsReleased:      isReleased,
		Language:        language,
		LastModified:    lastModified,
		Name:            name,
		Overview:        overview,
		Path:            path,
		PageURL:         pageURL,
		PresentedBy:     presentedBy,
		Price:           price,
		ReleaseDateMask: releaseDateMask,
		Series:          series,
		Slug:            slug,
		TagID:           tagID,
		Timestamp:       timestamp,
		Type:            typeOf,
		TypeAlias:       typeAlias,
		Unix:            unix,
		UPC:             UPC,
		URI:             URI,
		URL:             URL}
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
