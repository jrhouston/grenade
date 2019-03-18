package webserver

import (
	"fmt"
	"net/http"
)

// Run the web server
func Run(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "BOOM")
	})

	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		// TODO break this
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		// TODO break this
		fmt.Fprintf(w, "OK")
	})

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
