package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func net(rawurl string) (*xhr, error) {
	var (
		b   *[]byte
		err error
		ok  bool
		req *http.Request
		res *http.Response

		x = &xhr{
			Status:     http.StatusText(http.StatusBadRequest),
			StatusCode: http.StatusBadRequest}
	)
	req, err = http.NewRequest(http.MethodGet, rawurl, nil)
	ok = (err == nil)
	if !ok {
		return x, err
	}
	res, err = (&http.Client{}).Do(req)
	ok = (err == nil)
	if !ok {
		return x, err
	}
	b, err = netRead(res)
	ok = (err == nil)
	if !ok {
		return x, err
	}
	ok = (b != nil)
	if !ok {
		return x, fmt.Errorf("*b is nil")
	}
	ok = (len(*b) != 0)
	if !ok {
		return x, fmt.Errorf("*b is empty")
	}
	x.Body = b
	return x, err
}

func netGoQuery(rawurl string) (*goquery.Document, error) {
	var (
		doc *goquery.Document
		err error
		ok  bool
		req *http.Request
		res *http.Response
	)
	ok = (strings.HasPrefix(rawurl, "https://") || strings.HasPrefix(rawurl, "http://"))
	if !ok {
		return nil, fmt.Errorf("%s is not a valid URL", rawurl)
	}
	req, err = http.NewRequest(http.MethodGet, rawurl, nil)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	res, err = (&http.Client{}).Do(req)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return nil, fmt.Errorf(res.Status)
	}
	doc, err = goquery.NewDocumentFromResponse(res)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return doc, nil
}

func netRead(r *http.Response) (*[]byte, error) {
	var (
		b   []byte
		err error
		ok  bool
	)
	b, err = ioutil.ReadAll(r.Body)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &b, err
}
