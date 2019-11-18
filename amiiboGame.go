package main

type amiiboGame struct {
	Image *image   `json:"image"`
	Name  string   `json:"name"`
	URL   *address `json:"url"`
}
