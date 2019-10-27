package main

import (
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
	err = os.Remove(p)
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
	return err
}
