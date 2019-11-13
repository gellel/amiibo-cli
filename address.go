package main

import (
	"net/url"
	"strings"
	"text/tabwriter"
)

var (
	_ valuer = (&address{})
)

type address struct {
	Fragment string `json:"fragment"`
	Host     string `json:"host"`
	Hostname string `json:"hostname"`
	Path     string `json:"path"`
	Scheme   string `json:"scheme"`
	TLD      string `json:"tld"`
	URL      string `json:"url"`
}

func (a *address) Value() interface{} {
	return *a
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
		return nil, err
	}
	a = &address{
		Fragment: URL.Fragment,
		Host:     URL.Host,
		Hostname: URL.Hostname(),
		Path:     URL.Path,
		Scheme:   URL.Scheme,
		TLD:      URL.Host[:strings.IndexByte(URL.Host, '.')],
		URL:      rawurl}
	return a, err
}

func tableAddress(w *tabwriter.Writer, a *address) error {
	return printlnTable(w, *a)
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

func writeAddress(path, folder string, a *address) error {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalAddress(a)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, a.Path, b)
}
