package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type gameAmiibo struct {
	Image      *image    `json:"image"`
	IsReleased bool      `json:"is_released"`
	Name       string    `json:"name"`
	Series     string    `json:"series"`
	Timestamp  time.Time `json:"timestamp"`
	URL        *address  `json:"url"`
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
		template = "%s%s"
	)
	var (
		game       gameAmiibo
		image, _   = newImage(fmt.Sprintf(template, nintendoURL, s.Find("img").First().Text()))
		isReleased bool
		name       = stripAmiiboName(s.Find(".amiibo-name").Text())
		series     = strings.TrimSpace(s.Find(".isRelatedTo").Text())
		timestamp  time.Time
		URL        *address
	)
	game = gameAmiibo{
		Image:      image,
		IsReleased: isReleased,
		Name:       name,
		Series:     series,
		Timestamp:  timestamp,
		URL:        URL}
	return &game, nil
}
