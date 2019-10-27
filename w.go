package main

import "os"

func hasDir(fullpath string) bool {
	return (hasNotDir(fullpath) == false)
}

func hasNotDir(fullpath string) bool {
	var (
		err error
	)
	_, err = os.Stat(fullpath)
	return os.IsNotExist(err)
}

func makeDir(fullpath, folder string) {}
