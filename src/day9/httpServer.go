package main

import (
	"net/http"
	"fmt"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe(":8080", nil)
}
