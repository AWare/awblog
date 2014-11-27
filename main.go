package main

import (
	"HTML/template"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/AWare/awblog/post"
	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", renderPost())
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
func renderIndex(rw http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("markdown/")
	if err != nil {
		log.Fatal("FILE ERROR")
	}
	for _, f := range files {
		fmt.Fprintf(rw, f.Name(), html.EscapeString(r.URL.Path))
	}
}
func renderPost() func(rw http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		a := path.Clean(r.URL.String())
		fmt.Println(a)
		if a == "/" {
			renderIndex(rw, r)
			return
		}
		p, err := post.NewPostFromFile("markdown/" + a)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		if err := t.Execute(rw, p); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}
