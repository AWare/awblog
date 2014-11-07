package main

import (
	"HTML/template"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

type Post struct {
	Title   string
	Image   bool
	Content template.HTML
}

func (p *Post) ReadPostFromFile(name string) error {
	f, err := ioutil.ReadFile("markdown" + name + ".md")
	if err != nil {
		log.Println("file not found")
		//rw.WriteHeader(404)
		//fmt.Fprintf(rw, "Not Found.")
		return err
	}
	hasImage := true
	_, err = os.Open("markdown" + name + ".jpg")
	if err != nil {
		hasImage = false
	}
	p.Title = name[1:] //Check that name should still have the / removed.
	p.Content = template.HTML(blackfriday.MarkdownCommon(f))
	p.Image = hasImage
	return nil
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
		p := new(Post)
		p.ReadPostFromFile(a)
		if err := t.Execute(rw, p); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}
