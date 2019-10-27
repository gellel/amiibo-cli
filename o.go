package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func caller() (string, error) {
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

func userHomeDir() (string, error) {
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
		return s, fmt.Errorf("cannot find home directory for user")
	}
	return s, err
}
