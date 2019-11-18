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
	Domain    string `json:"domain"`
	Fragment  string `json:"fragment"`
	Host      string `json:"host"`
	Hostname  string `json:"hostname"`
	Path      string `json:"path"`
	Scheme    string `json:"scheme"`
	Subdomain string `json:"subdomain"`
	TLD       string `json:"tld"`
	URL       string `json:"url"`
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
	var (
		hostname = URL.Hostname()
	)
	var (
		domain    = parseAddressDomain(hostname)
		subdomain = parseAddressSubdomain(hostname)
		TLD       = parseAddressTLD(subdomain, domain, hostname)
	)
	a = &address{
		Domain:    domain,
		Fragment:  URL.Fragment,
		Host:      URL.Host,
		Hostname:  hostname,
		Path:      URL.Path,
		Scheme:    URL.Scheme,
		Subdomain: subdomain,
		TLD:       TLD,
		URL:       rawurl}
	return a, err
}

func parseAddressDomain(hostname string) string {
	const (
		sep string = "."
	)
	var (
		n          int
		ok         bool
		substrings = strings.Split(hostname, sep)
	)
	n = len(substrings)
	ok = (n < 4)
	if ok {
		return substrings[n-2]
	}
	return substrings[n-3]
}

func parseAddressSubdomain(hostname string) string {
	const (
		sep string = "."
	)
	var (
		n          int
		ok         bool
		subdomain  string
		substrings = strings.Split(hostname, sep)
	)
	n = len(substrings)
	ok = (n < 4)
	if ok {
		n = (n - 3)
	} else {
		n = (n - 4)
	}
	subdomain = substrings[n]
	return subdomain
}

func parseAddressTLD(subdomain, domain, hostname string) string {
	const (
		sep string = "."
	)
	var (
		TLD string
	)
	TLD = strings.Replace(hostname, (subdomain + sep), "", 1)
	TLD = strings.Replace(TLD, (domain + sep), "", 1)
	return TLD
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
