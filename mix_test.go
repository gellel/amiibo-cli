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

	x, err := newMixGameMapFromMix(m)
	if err != nil {
		panic(err)
	}
	//for k, v := range *x {
	//	fmt.Println(k, "->", v.compatabilityGame != nil, v.compatabilityItem != nil)
	//}
	y, err := newGameMap(x)
	if err != nil {
		panic(err)
	}
	for k, v := range *y {
		//fmt.Println(k, "->", stringifyMarshal(v))
		fmt.Println(k, "->", "\t", v.Timestamp)
	}
	a, err := newMixAmiiboMapFromMix(m)
	if err != nil {
		panic(err)
	}
	z, err := newAmiiboMap(a)
	if err != nil {
		panic(err)
	}
	for k, v := range *z {
		fmt.Println(k, "->", "\t", v.Timestamp)

	}
	/*
		x, err := newMixAmiiboMapFromMix(m)
		if err != nil {
			t.Fatalf(err.Error())
		}
		z, _ := newAmiiboMap(x)
		for _, a := range *z {
			fmt.Println(stringifyMarshal(a))
		}
	*/
	//newMixItemMapFromMix(m)
	/*
		for k, v := range *x {
			fmt.Println(k, "\t", v.compatabilityAmiibo != nil, v.lineupAmiibo != nil, v.lineupItem != nil)
		}
	*/
	/*
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

		for _, v := range m.CompatabilityGame {
			fmt.Println(v.Path)
			fmt.Println(v.URL)
		}
		for _, v := range m.CompatabilityItem {
			fmt.Println(v.Path)
			fmt.Println(v.URL)
		}
	*/
}
