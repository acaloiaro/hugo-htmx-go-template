package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:public
var content embed.FS

func main() {
	mux := http.NewServeMux()
	serverRoot, _ := fs.Sub(content, "public")

	// Serve all hugo content (the 'public' directory) at the root url
	mux.Handle("/", http.FileServer(http.FS(serverRoot)))

	// Add any number of handlers for custom endpoints here
	mux.HandleFunc("/hello_world", helloWorld)
	mux.HandleFunc("/hello_world_form", helloWorldForm)

	fmt.Printf("Starting API server on port 1314\n")
	if err := http.ListenAndServe("0.0.0.0:1314", mux); err != nil {
		log.Fatal(err)
	}
}

// the handler accepts GET requests to /hello_world
// It checks the URL params for the "name" param and populates the html/template variable with its value
// if no "name" url parameter is present, "name" is defaulted to "World"
//
// It responds with the the HTML partial `partials/helloworld.html`
func helloWorld(w http.ResponseWriter, r *http.Request) {
	// in development, the Origin is the the Hugo server, i.e. http://localhost:1313
	// but in production, it is the domain name where one's site is deployed
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "null" || name == "" {
		name = "World"
	}

	tmpl := template.Must(template.ParseFiles("partials/helloworld.html"))
	var buff = bytes.NewBufferString("")
	err := tmpl.Execute(buff, map[string]string{"Name": name})
	if err != nil {
		ise(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buff.Bytes())
}

// this handler accepts POST requests to /hello_world_form
// It checks the post request body for the form value "name" and populates the html/template
// variable with its value
//
// It responds with a simple greeting HTML partial
func helloWorldForm(w http.ResponseWriter, r *http.Request) {
	// in development, the Origin is the the Hugo server, i.e. http://localhost:1313
	// but in production, it is the domain name where one's site is deployed
	// for this demo, we're using `*` to keep things simple
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, hx-current-url, hx-request")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	name := "World"
	// The name is not in the query param, let's see if it was submitted as a form
	if err := r.ParseForm(); err != nil {
		ise(err, w)
		return
	}

	name = r.FormValue("name")
	// we're dealing with a really simple template; let's use an inline string instead of a whole separate file for this
	// one
	tmpl, err := template.New("form_response").Parse("<h3>Greeting: Hello, {{ .Name }}!</h3>")
	if err != nil {
		ise(err, w)
		return
	}

	var buff = bytes.NewBufferString("")
	err = tmpl.Execute(buff, map[string]string{"Name": name})
	if err != nil {
		ise(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buff.Bytes())
}

func ise(err error, w http.ResponseWriter) {
	fmt.Fprintf(w, "error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
}
