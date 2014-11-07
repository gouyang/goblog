package blog

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/ouyanggh/goblog/core"
	"github.com/ouyanggh/goblog/models"
)

var oldtitle string

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, p *models.Post, tmpl string) {
	t, _ := template.ParseFiles("./templates/" + tmpl + ".html")
	t.Execute(w, p)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/layout.html")
	t.Execute(w, "")
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/new.html")
	t.Execute(w, "")
}

func SavePost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	now := time.Now().UTC()
	p := &models.Post{Title: title, Created: now, Body: []byte(body)}
	core.SqliteInsert(p)
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
	//http.Redirect(w, r, "/blogs", http.StatusFound)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/update/"):]
	p := core.SqliteQuery(title)
	//p := &models.Post{Title: title, Body: []byte(body)}
	oldtitle = title
	renderTemplate(w, p, "edit")
}

func SaveUpdate(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	now := time.Now().UTC()
	p := &models.Post{Title: title, Created: now, Body: []byte(body)}
	//p := &models.Post{Title: title, Body: []byte(body)}
	core.SqliteUpdate(p, oldtitle)
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/delete/"):]
	core.SqliteDelete(title)
	http.Redirect(w, r, "/blogs/delete/", http.StatusFound)
}

func PostsForDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/\">Home</a>")
	fmt.Fprintf(w, "<h1>Blog posts ...</h1>")
	titles := core.SqliteQueryAllPost()
	for title, _ := range titles {
		fmt.Fprintf(w, "<div><strong><em><a href=\"/blog/%s\">%s</a></em></strong></div><div><a href=\"/blog/delete/%s\">delete</a></div></br>", title, title, title)
	}
	title := r.URL.Path[len("/blog/"):]
	core.SqliteDelete(title)
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/"):]
	p := core.SqliteQuery(title)
	//p := &models.Post{Title: title, Body: []byte(body)}
	renderTemplate(w, p, "view")
}

func ListPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<head>")
	fmt.Fprintf(w, "<link rel=\"stylesheet\" href=\"/static/css/div.css\">")
	fmt.Fprintf(w, "</head>")
	fmt.Fprintf(w, "<a href=\"/\">Home</a>")
	fmt.Fprintf(w, "<h1>Blog posts ...</h1>")
	titles := core.SqliteQueryAllPost()
	for title, _ := range titles {
		fmt.Fprintf(w, "<div><strong><em><a href=\"/blog/%s\">%s</a></em></strong></div>", title, title)
		fmt.Fprintf(w, "</br>")
		fmt.Fprintf(w, "Created on %s</br></br><div class=\"content\">%s</div>", titles[title].Created, titles[title].Body)
		fmt.Fprintf(w, "</br></br></br>")
	}
	fmt.Fprintf(w, "</html>")
}
