package main

import (
	"log"
	"net/http"
	"text/template"
)

// Temps is a struct that holds all the templates.
type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></bod></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

func setupTemp() *Temps {
	temps := new(Temps)
	temps.notemp = notemp()

	// set index template
	indx, er := template.ParseFiles("templates/index.html")
	if er != nil {
		indx = temps.notemp
	}
	temps.indx = indx

	// set hello template
	helo, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		helo = temps.notemp
	}
	temps.helo = helo

	return temps
}

// index handler
func index(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler
func hello(w http.ResponseWriter, rq *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

// main program
func main() {
	temps := setupTemp()
	// index handling
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq, temps.indx)
	})
	// hello handling
	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq, temps.helo)
	})

	http.ListenAndServe("", nil)
}
