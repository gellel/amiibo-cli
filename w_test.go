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
	err = makeDir(d, "amiibo-cli-test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = delDir(d, "amiibo-cli-test")
	if err != nil {
		t.Fatalf(err.Error())
	}
}
