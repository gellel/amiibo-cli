package main

import (
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"golang.org/x/text/transform"
)

func normalizeURI(s string) string {
	s, _, _ = transform.String(transformer, s)
	s = replacerURI.Replace(s)
	s = regexpUnwantedURI.ReplaceAllString(s, "-")
	s = regexpHyphens.ReplaceAllString(s, "")
	s = strings.TrimSuffix(s, "-")
	s = strings.ToLower(s)
	return s
}

func normalizeAmiiboMapKey(s string) string {
	const (
		p1 string = "/content/noa/en_US"
		p2 string = "/amiibo/detail/"
		s1 string = ".html"
	)
	s = strings.TrimPrefix(s, p1)
	s = strings.TrimPrefix(s, p2)
	s = strings.TrimSuffix(s, s1)
	return s
}

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
		switch reflect.ValueOf(value).Kind() == reflect.Ptr {
		case true:
			printlnTable(w, value.(valuer).Value())
		default:
			fmt.Fprintln(w, fmt.Sprintf(s, i, key, value))
		}
	}
	return w.Flush()
}

func fprintlnTable(w *tabwriter.Writer, v valuer) error {
	var (
		err error
	)
	for key, value := range mapMarshal(v) {
		_, err = fmt.Fprintln(w, fmt.Sprintf("t%s\t%v", key, value))
		if err != nil {
			break
		}
	}
	return err
}
