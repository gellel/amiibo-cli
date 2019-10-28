package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func writeFile(path, folder, name, ext string, b *[]byte) error {
	var (
		err error
		ok  bool
		p   = filepath.Join(path, folder, fmt.Sprintf("%s.%s", name, ext))
	)
	err = makeDir(path, folder)
	ok = (err == nil)
	if !ok {
		return err
	}
	return ioutil.WriteFile(p, *b, writeMode)
}

func writeJSON(path, folder, name string, b *[]byte) error {
	const (
		ext string = ".json"
	)
	var (
		ok bool
	)
	ok = strings.HasPrefix(name, ".")
	if ok {
		name = strings.TrimPrefix(name, ".")
	}
	ok = strings.HasSuffix(name, ext)
	if ok {
		name = strings.TrimSuffix(name, ext)
	}
	return writeFile(path, folder, name, ext, b)
}
