package main

import "fmt"

var (
	_ hashMap = (&byteMap{})
)

type byteMap map[string]*[]byte

func (m *byteMap) Del(key string) bool {
	delete((*m), key)
	return (m.Has(key) == false)
}

func (m *byteMap) Each(fn func(string, interface{})) {
	for key, value := range *m {
		fn(key, value)
	}
}

func (m *byteMap) Get(key string) (*[]byte, bool) {
	var (
		b  *[]byte
		ok bool
	)
	b, ok = ((*m)[key])
	return b, ok
}

func (m *byteMap) Has(key string) bool {
	var (
		ok bool
	)
	_, ok = m.Get(key)
	return ok
}

func (m *byteMap) Keys() []string {
	var (
		keys []string
	)
	m.Each(func(key string, _ interface{}) {
		keys = append(keys, key)
	})
	return keys
}

func (m *byteMap) Len() int {
	return len(*m)
}

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
