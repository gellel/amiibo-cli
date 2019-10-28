package main

import (
	"errors"
	"fmt"
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

const (
	writeMode os.FileMode = 0777
)

var (
	errBNil = errors.New("*[]byte is nil")
)

var (
	errBEmpty = errors.New("*[]byte is empty")
)

var (
	errNotDir = errors.New("p is not a directory")
)

var (
	errNotPtr = errors.New("interface{} is not a uintptr")
)

var (
	errNoHomeDir = fmt.Errorf("cannot find home directory for user")
)

var (
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
)

func main() {}
