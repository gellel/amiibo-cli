package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalB(b *[]byte, v interface{}) error {
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
