package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
	"unicode"

	"github.com/gorilla/mux"
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
	regexpHyphens = regexp.MustCompile(`\-{2,}`)
)

var (
	regexpUnwantedURI = regexp.MustCompile(`[^a-zA-Z0-9&]+`)
)

var (
	replacerURI = strings.NewReplacer([]string{"&", "and", "'", ""}...)
)

var (
	transformer = transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool { return unicode.Is(unicode.Mn, r) }), norm.NFC)
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	m, err := getMix()
	if err != nil {
		panic(err)
	}
	gm, err := newMixGameMapFromMix(m)
	if err != nil {
		panic(err)
	}
	games, err := newGameMap(gm)
	if err != nil {
		panic(err)
	}
	am, err := newMixAmiiboMapFromMix(m)
	if err != nil {
		panic(err)
	}
	amiibos, err := newAmiiboMap(am)
	if err != nil {
		panic(err)
	}
	s := amiibos.Values()
	as, err := marshal(&s)
	if err != nil {
		panic(err)
	}
	router.Handle("/amiibo", amiiboMuxAll{Amiibo: *as}).Methods(http.MethodGet)
	router.Handle("/games", gameMuxAll{Games: games.Values()}).Methods(http.MethodGet)
	router.Handle("/games/{name}", gameMuxName{Games: games}).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
