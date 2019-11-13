package main

import "fmt"

var (
	_ hashMap = (&amiiboMap{})
)

type amiiboMap map[string]*amiibo

func (m *amiiboMap) Del(key string) bool {
	delete((*m), key)
	return (m.Has(key) == false)
}

func (m *amiiboMap) Each(fn func(string, interface{})) {
	for key, value := range *m {
		fn(key, value)
	}
}

func (m *amiiboMap) Get(key string) (*amiibo, bool) {
	var (
		amiibo *amiibo
		ok     bool
	)
	amiibo, ok = ((*m)[key])
	return amiibo, ok
}

func (m *amiiboMap) Has(key string) bool {
	var (
		ok bool
	)
	_, ok = m.Get(key)
	return ok
}

func (m *amiiboMap) Keys() []string {
	var (
		keys []string
	)
	m.Each(func(key string, _ interface{}) {
		keys = append(keys, key)
	})
	return keys
}

func (m *amiiboMap) Len() int {
	return len(*m)
}

func (m *amiiboMap) Values() []*amiibo {
	var (
		amiibos []*amiibo
	)
	for _, amiibo := range *m {
		amiibos = append(amiibos, amiibo)
	}
	return amiibos
}

func newAmiiboMap(m *mixAmiiboMap) (*amiiboMap, error) {
	var (
		a   *amiibo
		err error
		ok  bool

		x = amiiboMap{}
	)
	ok = (m != nil)
	if !ok {
		return nil, fmt.Errorf("*m is nil")
	}
	ok = (len(*m) != 0)
	if !ok {
		return nil, fmt.Errorf("*m is empty")
	}
	for _, v := range *m {
		a, err = newAmiibo(v.compatabilityAmiibo, v.lineupAmiibo, v.lineupItem)
		if err != nil {
			continue
		}
		if a == nil {
			continue
		}
		x[a.URI] = a
	}
	return &x, err
}

func newAmiiboMapFromMix(m *mix) (*amiiboMap, error) {
	var (
		err          error
		mixAmiiboMap *mixAmiiboMap
		ok           bool
	)
	mixAmiiboMap, err = newMixAmiiboMapFromMix(m)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return newAmiiboMap(mixAmiiboMap)
}
