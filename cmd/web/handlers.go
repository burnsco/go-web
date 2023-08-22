package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	//  check to see if the url path is correct, if not then send notFound error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// html go template layout for my site
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// parsing the above files, checking for errors (code 500)
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific post with ID %d...", id)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	const NOT_FOUND_ERROR = 405

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")

		http.Error(w, "Method Not Allowed", NOT_FOUND_ERROR)
		return
	}

	w.Write([]byte("Create a new post..."))
}
