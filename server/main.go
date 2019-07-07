package main

import (
	"./cache"
	"./tcp"
	"./http"
)

func main() {
	ca := cache.New("inmemory")
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
