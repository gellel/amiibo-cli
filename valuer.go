package main

type valuer interface {
	Value() interface{}
}

type keyspace interface {
	Key() string
}
