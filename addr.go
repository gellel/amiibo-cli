package main

import "net/url"

type addr struct {
	Fragment string `json:"fragment"`
	Host     string `json:"host"`
	Hostname string `json:"hostname"`
	Path     string `json:"path"`
	Scheme   string `json:"scheme"`
	URL      string `json:"URL"`
}

func newAddr(rawurl string) (*addr, error) {
	var (
		a   *addr
		err error
		ok  bool
		URL *url.URL
	)
	URL, err = url.Parse(rawurl)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	a = &addr{
		Fragment: URL.Fragment,
		Host:     URL.Host,
		Hostname: URL.Hostname(),
		Path:     URL.Path,
		Scheme:   URL.Scheme,
		URL:      rawurl}
	return a, err
}
