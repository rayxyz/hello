package main

import (
	"net/http"
	_ "net/http/pprof"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	// http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", router)
}
