package main

import "fmt"

type amiiboMap map[string]*amiibo

func (a *amiiboMap) Values() []*amiibo {
	var (
		amiibos []*amiibo
	)
	for _, amiibo := range *a {
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
