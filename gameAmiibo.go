package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type gameAmiibo struct {
	Image           *image    `json:"image"`
	IsReleased      bool      `json:"is_released"`
	Name            string    `json:"name"`
	ReleaseDateMask string    `json:"release_date_mask"`
	Series          string    `json:"series"`
	Timestamp       time.Time `json:"timestamp"`
	URL             *address  `json:"url"`
}

func newGameAmiibo(s *goquery.Selection) (*gameAmiibo, error) {
	var (
		ok bool
	)
	ok = (s != nil)
	if !ok {
		return nil, fmt.Errorf("*s is nil")
	}
	ok = (s.Length() != 0)
	if !ok {
		return nil, fmt.Errorf("*s is empty")
	}
	const (
		template   string = "%s%s"
		timeLayout string = "01/02/2006"
	)
	var (
		game            gameAmiibo
		image, _        = parseGameAmiiboImage(s)
		isReleased      bool
		name, _         = parseGameAmiiboName(s)
		releaseDateMask string
		series, _       = parseGameAmiiboSeries(s)
		timestamp       time.Time
		URL, _          = parseGameAmiiboURL(s)
	)
	releaseDateMask, _ = parseGameAmiiboReleaseDateMask(s)
	timestamp, _ = time.Parse(timeLayout, releaseDateMask)
	timestamp = timestamp.UTC()
	isReleased = timestamp.Unix() < time.Now().UTC().Unix()
	game = gameAmiibo{
		Image:           image,
		IsReleased:      isReleased,
		Name:            name,
		ReleaseDateMask: releaseDateMask,
		Series:          series,
		Timestamp:       timestamp,
		URL:             URL}
	return &game, nil
}

func parseGameAmiiboImage(s *goquery.Selection) (*image, error) {
	const (
		CSS string = "img"
	)
	var (
		ok     bool
		rawurl string
	)
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return nil, fmt.Errorf("*s is empty")
	}
	rawurl, ok = s.Attr("src")
	if !ok {
		return nil, fmt.Errorf("*s has no src")
	}
	return newImage(fmt.Sprintf("%s%s", nintendoURL, rawurl))
}

func parseGameAmiiboName(s *goquery.Selection) (string, error) {
	const (
		CSS string = ".amiibo-name"
	)
	var (
		err  error
		name string
		ok   bool
	)
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return name, fmt.Errorf("*s is empty")
	}
	name = (s.Text())
	ok = (len(name) != 0)
	if !ok {
		return name, fmt.Errorf("*s has no text")
	}
	return stripAmiiboName(name), err
}

func parseGameAmiiboReleaseDateMask(s *goquery.Selection) (string, error) {
	const (
		CSS string = "span[itemprop='releaseDate']"
	)
	var (
		err       error
		ok        bool
		substring = (s.Find(CSS).Text())
	)
	substring = strings.TrimSpace(substring)
	ok = (len(substring) != 0)
	if !ok {
		return substring, fmt.Errorf("*s is empty")
	}
	substring = strings.ToLower(substring)
	substring = strings.Replace(substring, "available", "", 1)
	return substring, err
}

func parseGameAmiiboSeries(s *goquery.Selection) (string, error) {
	const (
		CSS string = "span[itemprop='isRelatedTo']"
	)
	var (
		err    error
		series string
		ok     bool
	)
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return series, fmt.Errorf("*s is empty")
	}
	series = (s.Text())
	ok = (len(series) != 0)
	if !ok {
		return series, fmt.Errorf("*s has no text")
	}
	series = strings.TrimSpace(series)
	return series, err
}

func parseGameAmiiboURL(s *goquery.Selection) (*address, error) {
	const (
		CSS string = "a"
	)
	var (
		ok     bool
		rawurl string
	)
	s = s.Find(CSS)
	ok = (s.Length() != 0)
	if !ok {
		return nil, fmt.Errorf("*s is empty")
	}
	rawurl, ok = s.Attr("href")
	if !ok {
		return nil, fmt.Errorf("*s has no href")
	}
	rawurl = fmt.Sprintf("%s%s", nintendoURL, rawurl)
	return newAddress(rawurl)
}
