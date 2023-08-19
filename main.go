package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from corey webserver"))
}

func main() {
	// initialize new server
	mux := http.NewServeMux()

	// pass home function to handle requests to the site "/"
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
