package main

import (
	"fmt"
	"reflect"
	"text/tabwriter"
)

func printlnTable(w *tabwriter.Writer, x interface{}) error {
	const (
		s string = "%d\t%s\t%v"
	)
	var (
		i int
		n int
		v reflect.Value
		t reflect.Type
	)
	v = reflect.ValueOf(x)
	t = v.Type()
	n = v.NumField()
	for i = 0; i < n; i++ {
		var (
			key   = t.Field(i).Name
			value = v.Field(i).Interface()
		)
		fmt.Fprintln(w, fmt.Sprintf(s, i, key, value))
	}
	return w.Flush()
}
