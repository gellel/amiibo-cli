package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
