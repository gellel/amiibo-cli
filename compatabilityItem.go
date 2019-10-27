package main

import (
	"fmt"
	"text/tabwriter"
)

type compatabilityItem struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

func tableCompatabilityItem(w *tabwriter.Writer, c *compatabilityItem) error {
	var (
		err error
		ok  bool

		x = []string{
			fmt.Sprintf("description\t%s", c.Description),
			fmt.Sprintf("last modified\t%d", c.LastModified),
			fmt.Sprintf("path\t%s", c.Path),
			fmt.Sprintf("title\t%s", c.Title),
			fmt.Sprintf("url\t%s", c.URL)}
	)
	for i, s := range x {
		_, err = fmt.Fprintln(w, fmt.Sprintf("%d\t%s", i, s))
		ok = (err == nil)
		if !ok {
			return err
		}
	}
	return w.Flush()
}
