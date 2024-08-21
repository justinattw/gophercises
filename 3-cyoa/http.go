package main

import (
	"net/http"
)

func serve() {
	var handler http.HandlerFunc

	http.ListenAndServe(":8080", handler)
}
