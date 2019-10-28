package main

type lineupItem struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

func marshalLineupItem(l *lineupItem) (*[]byte, error) {
	return marshalB(l)
}
