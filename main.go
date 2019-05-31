package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("[PLACEHOLDER}"))
	})

	fmt.Println("127.0.0.1:3000")

	http.ListenAndServe(":3000", nil)
}