package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func getMux() (*mux.Router, error) {
	var (
		err error
		mix *mix
		ok  bool
	)
	mix, err = getMix()
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return newMux(mix)
}

func newMux(m *mix) (*mux.Router, error) {
	const (
		routeAmiibo     string = "/amiibo"
		routeAmiiboName string = routeAmiibo + "/" + "{name}"
		routeGames      string = "/games"
		routeGamesName  string = routeGames + "/" + "{name}"
	)
	var (
		amiiboMap     *amiiboMap
		amiiboMuxAll  *amiiboMuxAll
		amiiboMuxName *amiiboMuxName
		gameMap       *gameMap
		gameMuxAll    *gameMuxAll
		gameMuxName   *gameMuxName
		ok            bool
		router        *mux.Router
		wg            sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var (
			err error
		)
		amiiboMap, err = newAmiiboMapFromMix(m)
		if err != nil {
			return
		}
		amiiboMuxAll, err = newAmiiboMuxAll(amiiboMap.Values())
		if err != nil {
			return
		}
		amiiboMuxName, _ = newAmiiboMuxName(amiiboMap)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var (
			err error
		)
		gameMap, err = newGameMapFromMix(m)
		if err != nil {
			return
		}
		gameMuxAll, err = newGameMuxAll(gameMap.Values())
		if err != nil {
			return
		}
		gameMuxName, err = newGameMuxName(gameMap)
	}()
	wg.Wait()
	ok = (amiiboMuxAll != nil) && (amiiboMuxName != nil)
	if !ok {
		return nil, fmt.Errorf("cannot create amiibo mux routes")
	}
	ok = (gameMuxAll != nil) && (gameMuxName != nil)
	if !ok {
		return nil, fmt.Errorf("cannot create game mux routes")
	}
	router = mux.NewRouter().StrictSlash(true)
	router.Handle(routeAmiibo, amiiboMuxAll).Methods(http.MethodGet)
	router.Handle(routeAmiiboName, amiiboMuxName).Methods(http.MethodGet)
	router.Handle(routeGames, gameMuxAll).Methods(http.MethodGet)
	router.Handle(routeGamesName, gameMuxName).Methods(http.MethodGet)
	return router, nil
}
