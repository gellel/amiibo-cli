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
	x := map[string]int{}
	for _, a := range m.CompatabilityAmiibo {
		if _, ok := x[a.URL]; !ok {
			x[a.URL] = 0
		}
		x[a.URL]++
	}
	for _, a := range m.LineupAmiibo {
		if _, ok := x[a.DetailsURL]; !ok {
			x[a.DetailsURL] = 0
		}
		x[a.DetailsURL]++
	}
	for k, v := range x {
		fmt.Println(k, "->", v)
	}
}
