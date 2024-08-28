package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var m sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		count++
		m.Unlock()
		_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			_, _ = fmt.Fprintf(w, "Header[%q]: %q\n", k, v)
		}
		_, _ = fmt.Fprintf(w, "Host = %q\n", r.Host)
		_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			_, _ = fmt.Fprintf(w, "Form[%q]: %q\n", k, v)
		}
	})
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		_, _ = fmt.Fprintf(w, "Count: %d", count)
		m.Unlock()
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
