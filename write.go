package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readFile(fullpath string) (*[]byte, error) {
	var (
		b   []byte
		err error
		ok  bool
	)
	b, err = ioutil.ReadFile(fullpath)
	ok = err != nil
	if !ok {
		return nil, err
	}
	return &b, err
}

func readFolder(fullpath string) {
	var (
		err   error
		files []string
	)
	err = filepath.Walk(fullpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

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
