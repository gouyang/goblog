package blog

import (
	"fmt"
	"github.com/ouyanggh/goblog/core"
	"html/template"
	"io/ioutil"
	"net/http"
	//"strings"
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
	//p.saveArticle()
	core.SqliteInsert(p.Title, p.Body)
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
}

func BlogUpdate(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/update/"):]
	newtitle := r.FormValue("title")
	body := r.FormValue("body")
	p := &Article{Title: newtitle, Body: []byte(body)}
	//p.saveArticle()
	core.SqliteUpdate(p.Title, p.Body, title)
	http.Redirect(w, r, "/blog/"+newtitle, http.StatusFound)
}

func BlogEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/edit/"):]
	titles := core.SqliteQuery()
	p := &Article{Title: title, Body: titles[title]}
	renderTemplate(w, p, "edit")
}

func BlogDelete(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/delete/"):]
	fmt.Println(title)
	core.SqliteDelete(title)
	http.Redirect(w, r, "/blogs/delete", http.StatusFound)
}

func BlogsDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/\">Home</a>")
	fmt.Fprintf(w, "<h1>Blog posts ...</h1>")
	titles := core.SqliteQuery()
	for title, _ := range titles {
		fmt.Fprintf(w, "<div><strong><em><a href=\"/blog/%s\">%s</a></em></strong></div><div><a href=\"/blog/delete/%s\">delete</a></div>", title, title, title)
	}
	title := r.URL.Path[len("/blog/"):]
	core.SqliteDelete(title)
}

func BlogView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/"):]
	titles := core.SqliteQuery()
	p := &Article{Title: title, Body: titles[title]}
	renderTemplate(w, p, "view")
}

func BlogList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/\">Home</a>")
	fmt.Fprintf(w, "<h1>Blog posts ...</h1>")
	titles := core.SqliteQuery()
	for title, _ := range titles {
		fmt.Fprintf(w, "<div><strong><em><a href=\"/blog/%s\">%s</a></em></strong></div>", title, title)
	}
	//files, _ := ioutil.ReadDir("./files")
	//for _, f := range files {
	//	title := strings.Replace(f.Name(), ".txt", "", 1)
	//	fmt.Fprintf(w, "<div><strong><em><a href=\"/blog/%s\">%s</a></em></strong></div>", title, title)
	//}
}
