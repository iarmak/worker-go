package main

import (
	"net/http"

	"github.com/syumai/workers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, Go-GO!"
		w.Write([]byte(msg))
	})
	workers.Serve(nil) // use http.DefaultServeMux
}
