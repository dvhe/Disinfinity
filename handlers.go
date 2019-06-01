package main

import (
	"net/http"
)

//Index : serves default main page or whatever
func Index(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./Website"))
}
