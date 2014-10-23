package main

import (
	"HTML/template"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/codegangsta/negroni"
	"github.com/russross/blackfriday"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", renderPost())
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")

}

type markdownPage struct {
	Title   string
	Content template.HTML
}

func renderIndex(rw http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("markdown/")
	if err != nil {
		log.Fatal("FILE ERROR")
	}
	for _, f := range files {
		fmt.Fprintf(rw, f.Name(), html.EscapeString(r.URL.Path))
		//rw.Fprintf(f.Name())
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
		f, err := ioutil.ReadFile("markdown" + a + ".md")
		if err != nil {
			log.Println("file not found")
			rw.WriteHeader(404)
			fmt.Fprintf(rw, "Not Found.")
			return
		}
		c := struct {
			Title   string
			Content template.HTML
		}{
			a[1:],
			template.HTML(blackfriday.MarkdownCommon(f)),
		}
		if err := t.Execute(rw, c); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}
