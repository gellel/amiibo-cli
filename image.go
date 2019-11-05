package main

import (
	"path/filepath"
	"strings"
)

type image struct {
	Address   *address `json:"address"`
	Extension string   `json:"extension"`
	Name      string   `json:"name"`
}

func newImage(rawurl string) (*image, error) {
	var (
		address *address
		err     error
		ext     string
		i       *image
		name    string
		ok      bool
	)
	address, err = newAddress(rawurl)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	ext = filepath.Ext(rawurl)
	ext = strings.TrimPrefix(ext, ".")
	name = filepath.Dir(rawurl)
	name = filepath.Base(name)
	i = &image{
		Address:   address,
		Extension: ext,
		Name:      name}
	return i, err
}
