package main

import (
	"log"

	"github.com/BurntSushi/xgb/xproto"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func reverseStack(data []xproto.Window) []xproto.Window {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}
