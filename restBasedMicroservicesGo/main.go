package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/helloWorld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}
