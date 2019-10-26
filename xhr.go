package main

type xhr struct {
	Body       *[]byte `json:"body"`
	Status     string  `json:"status"`
	StatusCode int     `json:"status_code"`
}
