package main

import (
	"encoding/json"
	"reflect"
)

func marshalB(v interface{}) (*[]byte, error) {
	var (
		b   []byte
		err error
		k   = reflect.ValueOf(v).Kind()
		ok  bool
	)
	ok = (k == reflect.Ptr)
	if !ok {
		return nil, errNotPtr
	}
	b, err = json.Marshal(v)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	ok = (len(b) != 0)
	if !ok {
		return nil, errBEmpty
	}
	return &b, err
}

func unmarshalB(b *[]byte, v interface{}) error {
	var (
		err error
		ok  bool
	)
	ok = (b != nil)
	if !ok {
		return errBNil
	}
	ok = (len(*b) > 0)
	if !ok {
		return errBEmpty
	}
	err = json.Unmarshal(*b, v)
	ok = (err == nil)
	if !ok {
		return err
	}
	return err
}
