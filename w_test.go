package main

import (
	"testing"
)

func TestW(t *testing.T) {

	var (
		d, err = userHomeDir()
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	makeDir(d, "test")
}
