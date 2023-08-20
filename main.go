package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from corey webserver"))
}

func showPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Display a specific snippet here."))
}

func showJSON(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sucessfully opened users.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name": "Corey"}`))

	defer jsonFile.Close()
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
	mux.HandleFunc("/showJSON", showJSON)

	log.Println("Starting server on :4002")
	err := http.ListenAndServe(":4002", mux)
	log.Fatal(err)
}
