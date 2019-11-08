package main

import (
	"fmt"
	"strconv"
)

type game struct {
	Complete        bool     `json:"complete"`
	Description     string   `json:"description"`
	GamePath        string   `json:"game_path"`
	GameURL         *address `json:"game_url"`
	ID              string   `json:"id"`
	Image           *image   `json:"image"`
	IsReleased      bool     `json:"is_released"`
	LastModified    int64    `json:"last_modified"`
	Path            string   `json:"path"`
	Name            string   `json:"name"`
	ReleaseDateMask string   `json:"release_date_mask"`
	Title           string   `json:"title"`
	Type            string   `json:"type"`
	URI             string   `json:"uri"`
	URL             *address `json:"url"`
}

func newGame(c *compatabilityGame, i *compatabilityItem) (*game, error) {
	var (
		ok bool
	)
	ok = (c != nil) || (i != nil)
	if !ok {
		return nil, fmt.Errorf("*c, *l and *i are nil")
	}
	var (
		complete        bool
		description     string
		g               *game
		gamePath        string
		gameURL         *address
		ID              string
		image           *image
		isReleased      bool
		lastModified    int64
		path            string
		name            string
		releaseDateMask string
		title           string
		typeOf          string
		URI             string
		URL             *address
	)
	complete = (c != nil) && (i != nil)
	if c != nil {
		gamePath = c.Path
		gameURL, _ = newAddress(fmt.Sprintf("%s%s", nintendoURL, c.URL))
		ID = c.ID
		image, _ = newImage(fmt.Sprintf("%s%s", nintendoURL, c.Image))
		isReleased, _ = strconv.ParseBool(c.IsReleased)
		name = c.Name
		path = c.Path
		releaseDateMask = c.ReleaseDateMask
		typeOf = c.Type
	}
	if i != nil {
		description = i.Description
		lastModified = i.LastModified
		path = i.Path
		title = i.Title
		URL, _ = newAddress(fmt.Sprintf("%s%s", nintendoURL, i.URL))
	}
	URI = normalizeURI(name)
	g = &game{
		Complete:        complete,
		Description:     description,
		GamePath:        gamePath,
		GameURL:         gameURL,
		ID:              ID,
		Image:           image,
		IsReleased:      isReleased,
		LastModified:    lastModified,
		Path:            path,
		Name:            name,
		ReleaseDateMask: releaseDateMask,
		Title:           title,
		Type:            typeOf,
		URI:             URI,
		URL:             URL}
	return g, nil
}
