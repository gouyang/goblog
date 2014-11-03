package blog

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Article struct {
	Title string
	Body  []byte
}

func (p *Article) saveArticle() error {
	filename := "./files/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadArticle(title string) (*Article, error) {
	filename := "./files/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Article{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, p *Article, tmpl string) {
	t, _ := template.ParseFiles("./templates/" + tmpl + ".html")
	t.Execute(w, p)
}

func DefaultView(w http.ResponseWriter, r *http.Request) {
	homePage := &Article{Title: "Home Page", Body: []byte("A blog system")}
	renderTemplate(w, homePage, "layout")
}

func BlogNew(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	p := &Article{Title: title, Body: []byte(body)}
	renderTemplate(w, p, "new")
}

func BlogSave(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/blog/save/"):]
	title := r.FormValue("title")
	body := r.FormValue("body")
	p := &Article{Title: title, Body: []byte(body)}
	p.saveArticle()
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
}

func BlogEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/edit/"):]
	p, err := loadArticle(title)
	if err != nil {
		p = &Article{Title: title}
	}
	renderTemplate(w, p, "edit")
}

func BlogList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/\">Home</a>")
	fmt.Fprintf(w, "<h1>Blog posts ...</h1>")
	files, _ := ioutil.ReadDir("./files")
	for _, f := range files {
		title := strings.Replace(f.Name(), ".txt", "", 1)
		fmt.Fprintf(w, "<div><a href=\"/blog/%s\">%s</a></div>", title, title)
	}
}

func BlogView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/"):]
	p, _ := loadArticle(title)
	renderTemplate(w, p, "view")
}
