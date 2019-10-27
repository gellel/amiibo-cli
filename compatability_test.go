package main

import (
	"testing"
)

func TestCompatability(t *testing.T) {
	c, err := getCompatability()
	if err != nil {
		panic(err)
	}
	tableCompatability(w, c)
	tableCompatabilityAmiibo(w, c.Amiibo[0])
	tableCompatabilityItem(w, c.Items[0])
}
