package main

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

type Message struct {
	Type     int    `json:"type"`
	Nickname string `json:"nickname"`
	Body     string `json:"body"`
}

func MarshalByGo(v interface{}) ([]byte, error) {
	if b, err := json.Marshal(v); err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

var jsonn = jsoniter.ConfigCompatibleWithStandardLibrary

func MarshalByIterator(v interface{}) ([]byte,error) {
	if b, err := jsonn.Marshal(v); err != nil {
		return nil, err
	} else {
		return b, nil
	}
}
