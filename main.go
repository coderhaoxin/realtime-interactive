package main

import (
	"net/http"

	"./lib"
)

func main() {
	http.HandleFunc("/", lib.Service)
	http.ListenAndServe(":3000", nil)
}
