package main

import (
	"errors"
	"os"
	"text/tabwriter"
)

const (
	amiiboDetailURL string = amiiboURL + "/" + "detail"
)

const (
	amiiboURL string = nintendoURL + "/" + "amiibo"
)

const (
	compatabilityURI string = contentURI + "/" + "compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json"
)

const (
	contentURI string = nintendoURL + "/" + "content/noa/en_US/amiibo"
)

const (
	gamesDetailURL string = gamesURL + "/" + "detail"
)

const (
	gamesURL string = amiiboURL + "/" + "games"
)

const (
	lineupURI string = contentURI + "/" + "/line-up/jcr:content/root/responsivegrid/lineup.model.json"
)

const (
	nintendo string = "nintendo"
)

const (
	nintendoTLD string = nintendo + ".com"
)

const (
	nintendoURL string = "https://" + nintendoTLD
)

var (
	errBNil = errors.New("*[]byte is nil")
)

var (
	errBEmpty = errors.New("*[]byte is empty")
)

var (
	errNotPtr = errors.New("interface{} is not a uintptr")
)

var (
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
)

func main() {}
