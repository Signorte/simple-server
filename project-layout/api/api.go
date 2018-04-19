package api

import (
	"fmt"
	"io"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/2", hello2Handler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("%v", r))
}

func hello2Handler(w http.ResponseWriter, r *http.Request) {
}
