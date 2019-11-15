package main

import "time"

type gameAmiibo struct {
	Image     *image `json:"image"`
	IsReleased bool `json:"is_released"`
	Name      string `json:"name"`
	Series    string `json:"series"`
	Timestamp time.Time `json:"timestamp"`
	URL       *address `json:"url"`
}
