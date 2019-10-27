package main

import (
	"testing"
)

func TestW(t *testing.T) {
	const (
		folder string = ".amiibo-cli-test"
	)
	var (
		d, err = userHomeDir()
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = makeDir(d, folder)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = makeFile(d, folder, "test", "txt", &[]byte{})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = delDirAll(d, folder)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
