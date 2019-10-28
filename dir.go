package main

import (
	"fmt"
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
		return fmt.Errorf("%s is not a dir", p)
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
		return fmt.Errorf("%s is not a dir", p)
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

func getCurrentDir() (string, error) {
	var (
		err error
		ok  bool
		s   string
	)
	s, err = os.Executable()
	ok = (err == nil)
	if !ok {
		return s, err
	}
	s = filepath.Dir(s)
	return s, err
}

func getHomeDir() (string, error) {
	var (
		err error
		s   string
		ok  bool
	)
	s, err = os.UserHomeDir()
	ok = (err == nil)
	if !ok {
		return s, err
	}
	ok = hasDir(s)
	if !ok {
		return s, fmt.Errorf("cannot get home dir")
	}
	return s, err
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
		return fmt.Errorf("%s is not a dir", p)
	}
	return err
}
