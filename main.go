package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from corey webserver"))
}

func showPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet here."))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	const methodNotAllowedError = 405

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// replaces w.WriteHeader(405) and w.Write([]byte("Method not allowed"))
		http.Error(w, "Method Not Allowed", methodNotAllowedError)
		return
	}
	w.Write([]byte("Create a new post"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/post", showPost)
	mux.HandleFunc("/post/create", createPost)

	log.Println("Starting server on :4002")
	err := http.ListenAndServe(":4002", mux)
	log.Fatal(err)
}
