package main

import (
	"net/http"
)

//Index : serves default main page or whatever
func Index(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("Website/Frontend/public"))
	fs.ServeHTTP(w,r)
}
