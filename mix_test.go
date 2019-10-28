package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMix(t *testing.T) {
	m, err := getMix()
	if err != nil {
		t.Fatalf(err.Error())
	}

	var mu sync.Mutex
	var wg sync.WaitGroup

	gameMap := map[string]*compatabilityGame{}

	aM := map[string]*mixAmiibo{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, g := range m.CompatabilityGame {
			_, ok := gameMap[g.URL]
			mu.Lock()
			if !ok {
				gameMap[g.URL] = g
			}
			if ok {
				fmt.Println(g.URL)
			}
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, x := range m.CompatabilityAmiibo {
			_, ok := aM[x.URL]
			mu.Lock()
			if !ok {
				aM[x.URL] = &mixAmiibo{Compatability: x}
			} else {
				aM[x.URL].Compatability = x
			}
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, x := range m.LineupAmiibo {
			_, ok := aM[x.DetailsURL]
			mu.Lock()
			if !ok {
				aM[x.DetailsURL] = &mixAmiibo{Lineup: x}
			} else {
				aM[x.DetailsURL].Lineup = x
			}
			mu.Unlock()
		}
	}()
	wg.Wait()

	for _, v := range gameMap {
		fmt.Println(v.Name, "->", v)
	}
	/*
		for _, v := range aM {
			tableCompatabilityAmiibo(w, v.Compatability)
			fmt.Println("-")
			tableLineupAmiibo(w, v.Lineup)
			fmt.Println("-")
			fmt.Println(gameMap[v.Compatability.IsRelatedTo])
			fmt.Println("=======")
		}

		/*
			for _, x := range m.CompatabilityAmiibo {
				x.URL
			}
			fmt.Println("-----")
			for _, x := range m.CompatabilityGame {
				fmt.Println("CompatabilityGame", x)
			}
			fmt.Println("-----")
			for _, x := range m.CompatabilityItem {
				fmt.Println("CompatabilityItem", x)
			}
			fmt.Println("-----")
			for _, x := range m.LineupAmiibo {
				fmt.Println("LineupAmiibo", x)
			}
			fmt.Println("-----")
			for _, x := range m.LineupItem {
				fmt.Println("LineupItem", x)
			}
	*/
}
