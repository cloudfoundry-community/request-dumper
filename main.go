package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	log.SetPrefix("request_dumper: ")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Output(1, fmt.Sprintf("INFO: Processing Request: %s %s %s", r.Proto, r.Method, r.URL))
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Output(1, fmt.Sprintf("ERROR: Failed to dump request: %s", err))
		}
		dump = append(dump, '\n')
		_, err = w.Write(dump)
		if err != nil {
			log.Output(1, fmt.Sprintf("ERROR: Failed to write response: %s", err))
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	hostStr := ":" + port

	log.Output(1, fmt.Sprintf("INFO: Listening on '%s'", hostStr))
	log.Fatal(http.ListenAndServe(hostStr, nil))
}
