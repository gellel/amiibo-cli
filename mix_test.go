package main

import (
	"fmt"
	"testing"
)

func TestMix(t *testing.T) {
	m, err := getMix()
	if err != nil {
		t.Fatalf(err.Error())
	}
	mu := intersectionMixAmiibo(m)

	mu.Range(func(k, i interface{}) bool {
		fmt.Println(k, i)
		return true
	})
}
