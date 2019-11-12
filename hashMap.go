package main

type hashMap interface {
	Del(string) bool
	Each(func(string, interface{}))
	Has(string) bool
	Keys() []string
	Len() int
}
