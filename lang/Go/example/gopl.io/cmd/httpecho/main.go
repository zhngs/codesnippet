package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", debug_handler)
	http.HandleFunc("/count", count_handler)
	log.Fatal(http.ListenAndServe("localhost:7878", nil))
}

func debug_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

var mu sync.Mutex
var count int

func count_handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "URL.Path = %q, count = %d\n", r.URL.Path, count)
	mu.Unlock()
}
