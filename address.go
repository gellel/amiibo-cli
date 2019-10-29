package main

import (
	"fmt"
	"net/url"
)

type address struct {
	Fragment string `json:"fragment"`
	Host     string `json:"host"`
	Hostname string `json:"hostname"`
	Path     string `json:"path"`
	Scheme   string `json:"scheme"`
	URL      string `json:"url"`
}

func marshalAddress(a *address) (*[]byte, error) {
	return marshal(a)
}

func newAddress(rawurl string) (*address, error) {
	var (
		a   *address
		err error
		ok  bool
		URL *url.URL
	)
	URL, err = url.Parse(rawurl)
	ok = (err == nil)
	if !ok {
		fmt.Println(rawurl)

		return nil, err
	}
	a = &address{
		Fragment: URL.Fragment,
		Host:     URL.Host,
		Hostname: URL.Hostname(),
		Path:     URL.Path,
		Scheme:   URL.Scheme,
		URL:      rawurl}
	return a, err
}

func unmarshalAddress(b *[]byte) (*address, error) {
	var (
		a   address
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
