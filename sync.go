package main

import "sync"

func syncMuKeyCount(mu *sync.Map, k string) {
	if _, ok := mu.Load(k); !ok {
		mu.Store(k, 0)
	}
	n, _ := mu.Load(k)
	mu.Store(k, n.(int)+1)
}
