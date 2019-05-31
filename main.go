package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./public"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
