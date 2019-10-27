package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func delDir(path string, folder string) error {
	var (
		err error
		ok  bool
		p   = filepath.Join(path, folder)
	)
	ok = hasNotDir(p)
	if ok {
		return err
	}
	ok = isDir(p)
	if !ok {
		return errNotDir
	}
	err = os.Remove(p)
	ok = (err == nil)
	if !ok {
		return err
	}
	return err
}

func delDirAll(path string, folder string) error {
	var (
		err error
		ok  bool
		p   = filepath.Join(path, folder)
	)
	ok = hasNotDir(p)
	if ok {
		return err
	}
	ok = isDir(p)
	if !ok {
		return errNotDir
	}
	err = os.RemoveAll(p)
	ok = (err == nil)
	if !ok {
		return err
	}
	return err
}

func hasDir(path string) bool {
	return (hasNotDir(path) == false)
}

func hasNotDir(path string) bool {
	var (
		err error
	)
	_, err = os.Stat(path)
	return os.IsNotExist(err)
}

func isDir(path string) bool {
	var (
		err  error
		info os.FileInfo
		ok   bool
	)
	info, err = os.Stat(path)
	ok = (err == nil)
	if !ok {
		return false
	}
	return info.IsDir()
}

func makeDir(path string, folder string) error {
	var (
		err error
		ok  bool
		p   = filepath.Join(path, folder)
	)
	ok = hasDir(p)
	if ok {
		return err
	}
	err = os.MkdirAll(p, writeMode)
	ok = (err == nil)
	if !ok {
		return err
	}
	ok = isDir(p)
	if !ok {
		return errNotDir
	}
	return err
}

func makeFile(path, folder, name, ext string, b *[]byte) error {
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
