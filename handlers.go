package main

import (
	"fmt"
	"net/http"
)

//Index : serves default main page or whatever
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
