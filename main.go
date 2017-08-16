package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("public/index.html")))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
