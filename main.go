package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./public"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", handler)
}
