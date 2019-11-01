package main

import (
	"testing"
)

func TestMix(t *testing.T) {
	m, err := getMix()
	if err != nil {
		t.Fatalf(err.Error())
	}

	c, err := newCompatabilityAmiiboMap(m.CompatabilityAmiibo)
	if err != nil {
		t.Fatalf(err.Error())
	}
	l, err := newLineupAmiiboMap(m.LineupAmiibo)
	if err != nil {
		t.Fatalf(err.Error())
	}
	x, err := newMixAmiiboMap(c, l)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, v := range *x {
		_, err := newAmiibo(v.compatabilityAmiibo, v.lineupAmiibo)
		if err != nil {
			break
		}
	}
}
