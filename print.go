package main

import (
	"fmt"
	"text/tabwriter"
)

func printlnTable(w *tabwriter.Writer, rows *[]string) error {
	var (
		err error
		ok  bool
	)
	for i, s := range *rows {
		_, err = fmt.Fprintln(w, fmt.Sprintf("%d\t%s", i, s))
		ok = (err == nil)
		if !ok {
			return err
		}
	}
	return w.Flush()
}
