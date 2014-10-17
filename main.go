package main

import (
	"HTML/template"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/codegangsta/negroni"
	"github.com/russross/blackfriday"
)

func main() {
	b, err := ioutil.ReadFile("file.md")
	if err != nil {
		log.Fatalln("Ooops")
	}
	a := blackfriday.MarkdownCommon(b)
	fmt.Println(string(a))

	mux := http.NewServeMux()

	mux.HandleFunc("/", markdownServe)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")

}

type markdownPage struct {
	Title   string
	Content template.HTML
}

func markdownServe(rw http.ResponseWriter, r *http.Request) {
	//fmt.Println(path.Split(r.URL.String())
	t, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	a := path.Clean(r.URL.String())
	fmt.Println(a)
	f, err := ioutil.ReadFile("markdown" + a + ".md")
	if err != nil {
		log.Println("file not found")
		rw.WriteHeader(404)
		fmt.Fprintf(rw, "Not Found.")
		return
	}
	c := markdownPage{Title: a[1:], Content: template.HTML(blackfriday.MarkdownCommon(f))}
	if err := t.Execute(rw, c); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Write(blackfriday.MarkdownCommon(f))
}
