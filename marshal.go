package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func marshal(v interface{}) (*[]byte, error) {
	var (
		b   []byte
		err error
		k   = reflect.ValueOf(v).Kind()
		ok  bool
	)
	ok = (k == reflect.Ptr)
	if !ok {
		return nil, fmt.Errorf("v is not pointer")
	}
	b, err = json.Marshal(v)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	ok = (len(b) != 0)
	if !ok {
		return nil, fmt.Errorf("*b is empty")
	}
	return &b, err
}

func stringifyMarshal(v interface{}) string {
	var (
		b   *[]byte
		err error
		s   string
		ok  bool
	)
	b, err = marshal(v)
	ok = (err == nil)
	if !ok {
		return s
	}
	return string(*b)
}

func unmarshal(b *[]byte, v interface{}) error {
	var (
		err error
		ok  bool
	)
	ok = (b != nil)
	if !ok {
		return fmt.Errorf("*b is nil")
	}
	ok = (len(*b) > 0)
	if !ok {
		return fmt.Errorf("*b is empty")
	}
	err = json.Unmarshal(*b, v)
	ok = (err == nil)
	if !ok {
		return err
	}
	return err
}
