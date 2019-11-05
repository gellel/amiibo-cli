package main

import (
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	lineupURI string = contentURI + "/" + "line-up/jcr:content/root/responsivegrid/lineup.model.json"
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
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
)

var (
	regexpHTML = regexp.MustCompile(`(<[^>]*>|\n(\s{1,})?)`)
)

var (
	regexpName = regexp.MustCompile(`(\&\#[0-9]+\;|™|\(|\))`)
)

var (
	regexpSpaces = regexp.MustCompile(`\s{2,}`)
)

var (
	regexpUnderscore = regexp.MustCompile(`\_{2,}`)
)

var (
	transformer = transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool { return unicode.Is(unicode.Mn, r) }), norm.NFC)
)

func main() {
	fmt.Println(compatabilityURI)
	fmt.Println("-")
	fmt.Println(lineupURI)
}
