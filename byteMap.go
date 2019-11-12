package main

import "fmt"

type byteMap map[string]*[]byte

func newByteMap(m hashMap) (*byteMap, error) {
	var (
		b   *[]byte
		err error
		ok  bool
		x   = byteMap{}
	)
	ok = m.Len() != 0
	if !ok {
		return nil, fmt.Errorf("m is empty")
	}
	m.Each(func(key string, value interface{}) {
		b, err = marshal(value)
		ok = (err == nil)
		if ok {
			x[key] = b
		}
	})
	return &x, err
}
