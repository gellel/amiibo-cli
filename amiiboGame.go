package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type amiiboGame struct {
	Image *image   `json:"image"`
	Name  string   `json:"name"`
	URL   *address `json:"url"`
}

func newAmiiboGame(s *goquery.Selection) (*amiiboGame, error) {
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
	var (
		amiibo   amiiboGame
		image, _ = parseAmiiboGameImage(s)
		name, _  = parseAmiiboGameName(s)
		URL, _   = parseAmiiboGameURL(s)
	)
	amiibo = amiiboGame{
		Image: image,
		Name:  name,
		URL:   URL}
	return &amiibo, nil
}

func parseAmiiboGameImage(s *goquery.Selection) (*image, error) {
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

func parseAmiiboGameName(s *goquery.Selection) (string, error) {
	const (
		CSS string = "a[title]"
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
	name, ok = s.Attr("title")
	if !ok {
		return name, fmt.Errorf("*s has no title")
	}
	return stripAmiiboName(name), err
}

func parseAmiiboGameURL(s *goquery.Selection) (*address, error) {
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
