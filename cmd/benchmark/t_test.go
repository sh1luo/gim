package main

import (
	"log"
	"testing"
)

var m = Message{
	Type:     2,
	Nickname: "shiluo",
	Body:     "asdsadasdsad",
}

func BenchmarkMarshalByGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := MarshalByGo(&m); err != nil {
			log.Printf("err - ByGo:%s", err)
		}
	}
}

func BenchmarkMarshalByIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := MarshalByIterator(&m); err != nil {
			log.Printf("err - Iterator:%s", err)
		}
	}
}