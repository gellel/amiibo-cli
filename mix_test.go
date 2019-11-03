package main

import (
	"fmt"
	"testing"
)

func TestMix(t *testing.T) {
	m, err := getMix()
	if err != nil {
		t.Fatalf(err.Error())
	}
	//x := map[string]int{}

	a, _ := newCompatabilityAmiiboMap(m.CompatabilityAmiibo)

	b, _ := newLineupAmiiboMap(m.LineupAmiibo)

	c, _ := newLineupItemMap(m.LineupItem)

	type x struct {
		*compatabilityAmiibo
		*lineupAmiibo
		*lineupItem
	}

	z := map[string]*x{}

	d := 0
	e := 0
	f := 0

	for k, v := range *a {
		if _, ok := z[k]; !ok {
			z[k] = &x{}
		}
		x := z[k]
		x.compatabilityAmiibo = v
		fmt.Println(k)
		if d == 0 {
			//fmt.Println(k)
			//fmt.Println(v.URL)
			d++
		}
	}
	fmt.Println("--------DONE---------")
	for k, v := range *b {
		if _, ok := z[k]; !ok {
			z[k] = &x{}
		}
		x := z[k]
		x.lineupAmiibo = v
		fmt.Println(k)

		if e == 0 {
			//fmt.Println(k)
			//fmt.Println(v.AmiiboPage)
			//fmt.Println(v.DetailsURL)
			e++
		}
	}
	fmt.Println("--------DONE---------")
	for k, v := range *c {
		if _, ok := z[k]; !ok {
			z[k] = &x{}
		}
		x := z[k]
		x.lineupItem = v
		fmt.Println(k)
		if f == 0 {
			//fmt.Println(k)
			//fmt.Println(v.URL)
			f++
		}
	}
	fmt.Println("--------DONE---------")

	for _, v := range z {
		fmt.Println(v)
	}

	/*
		for _, v := range m.CompatabilityAmiibo {
			if _, ok := x[v.Name]; !ok {
				x[v.Name] = 0
			}
			x[v.Name] = x[v.Name] + 1
		}

			for _, v := range m.CompatabilityItem {
				if _, ok := x[v.Title]; !ok {
					fmt.Println(v.Title)
					x[v.Title] = 0
				}
				x[v.Title] = x[v.Title] + 1
			}

		for _, v := range m.LineupAmiibo {
			if _, ok := x[v.AmiiboName]; !ok {
				x[v.AmiiboName] = 0
			}
			x[v.AmiiboName] = x[v.AmiiboName] + 1
		}
		for _, v := range m.LineupItem {
			if _, ok := x[v.Title]; !ok {
				x[v.Title] = 0
			}
			x[v.Title] = x[v.Title] + 1
		}
		for k, v := range x {
			fmt.Println(k, "\t", "->", v)
		}
	*/
}
