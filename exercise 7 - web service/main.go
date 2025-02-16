package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil) //localhost is default
}

func Handler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./menu.txt")
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}

	io.Copy(w, file)
}
