package main

import (
	"fmt"
	"text/tabwriter"
)

type compatabilityGame struct {
	Image           string `json:"image"`
	ID              string `json:"id"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	ReleaseDateMask string `json:"releaseDateMask"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}

func tableCompatabilityGame(w *tabwriter.Writer, c *compatabilityGame) error {
	var (
		x = &[]string{
			fmt.Sprintf("image\t%s", c.Image),
			fmt.Sprintf("id\t%s", c.ID),
			fmt.Sprintf("is released\t%s", c.IsReleased),
			fmt.Sprintf("name\t%s", c.Name),
			fmt.Sprintf("release date mask\t%s", c.ReleaseDateMask),
			fmt.Sprintf("type\t%s", c.Type),
			fmt.Sprintf("url\t%s", c.URL)}
	)
	return printlnTable(w, x)
}
