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
		x = &[]string{
			fmt.Sprintf("description\t%s", c.Description),
			fmt.Sprintf("last modified\t%d", c.LastModified),
			fmt.Sprintf("path\t%s", c.Path),
			fmt.Sprintf("title\t%s", c.Title),
			fmt.Sprintf("url\t%s", c.URL)}
	)
	return printlnTable(w, x)
}
