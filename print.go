package main

import (
	"fmt"
	"reflect"
	"text/tabwriter"
)

func printlnTable(w *tabwriter.Writer, x interface{}) error {
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
		fmt.Fprintln(w, fmt.Sprintf("%d\t%s\t%v", i, t.Field(i).Name, v.Field(i).Interface()))
	}
	return w.Flush()
}
