package main

import (
	"testing"
)

func TestLineup(t *testing.T) {
	var (
		l, err = getLineup()
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	tableLineup(w, l)
}
