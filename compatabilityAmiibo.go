package main

import (
	"fmt"
	"text/tabwriter"
)

type compatabilityAmiibo struct {
	ID              string `json:"id"`
	Image           string `json:"image"`
	IsRelatedTo     string `json:"isRelatedTo"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	ReleaseDateMask string `json:"releaseDateMask"`
	TagID           string `json:"tagid"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}

func marshalCompatabilityAmiibo(c *compatabilityAmiibo) (*[]byte, error) {
	return marshal(c)
}

func stringifyMarshalCompatabilityAmiibo(c *compatabilityAmiibo) string {
	return stringifyMarshal(c)
}

func tableCompatabilityAmiibo(w *tabwriter.Writer, c *compatabilityAmiibo) error {
	var (
		x = &[]string{
			fmt.Sprintf("id\t%s", c.ID),
			fmt.Sprintf("image\t%s", c.Image),
			fmt.Sprintf("is related to\t%s", c.IsRelatedTo),
			fmt.Sprintf("is released\t%s", c.IsReleased),
			fmt.Sprintf("name\t%s", c.Name),
			fmt.Sprintf("release date mask\t%s", c.ReleaseDateMask),
			fmt.Sprintf("tag id\t%s", c.ReleaseDateMask),
			fmt.Sprintf("type\t%s", c.Type),
			fmt.Sprintf("url\t%s", c.URL)}
	)
	return printlnTable(w, x)
}

func unmarshalCompatabilityAmiibo(b *[]byte) (*compatabilityAmiibo, error) {
	var (
		c   compatabilityAmiibo
		err error
		ok  bool
	)
	err = unmarshal(b, &c)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &c, err
}
