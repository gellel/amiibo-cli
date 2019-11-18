package main

import (
	"path/filepath"
	"strings"
)

type image struct {
	Address   *address `json:"address"`
	Dir       string   `json:"dir"`
	Extension string   `json:"extension"`
	Name      string   `json:"name"`
}

func newImage(rawurl string) (*image, error) {
	const (
		sep string = "."
	)
	var (
		address *address
		err     error
		ok      bool
	)
	address, err = newAddress(rawurl)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	var (
		dir  = filepath.Dir(rawurl)
		ext  = filepath.Ext(rawurl)
		i    *image
		name = filepath.Base(rawurl)
	)
	i = &image{
		Address:   address,
		Dir:       dir,
		Extension: strings.TrimPrefix(ext, sep),
		Name:      strings.TrimSuffix(name, ext)}
	return i, err
}
