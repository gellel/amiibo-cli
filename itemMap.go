package main

import "fmt"

type itemMap map[string]*item

func newItemMap(m *mixItemMap) (*itemMap, error) {
	var (
		err error
		i   *item
		ok  bool

		x = itemMap{}
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
		i, err = newItem(v.compatabilityItem, v.lineupItem)
		if err != nil {
			continue
		}
		fmt.Println(i)
	}
	return &x, err
}
