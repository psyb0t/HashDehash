package main

import (
	"HashDehash/handlers"
	"log"
	"net/http"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/hash", handlers.Hash)
	router.HandleFunc("/hash/", handlers.Hash)

	router.HandleFunc("/dehash", handlers.Dehash)
	router.HandleFunc("/dehash/", handlers.Dehash)

	log.Fatal(http.ListenAndServe("127.0.0.1:8050", router))
}
