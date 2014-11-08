package blog

import (
	"html/template"
	"log"
	"net/http"
	"path"
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
	rtmpl := path.Join("templates", tmpl+".html")
	t, err := template.ParseFiles(rtmpl)
	CheckErr(err)
	err = t.Execute(w, p)
	CheckErr(err)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := path.Join("templates", "layout.html")
	t, err := template.ParseFiles(tmpl)
	CheckErr(err)
	err = t.Execute(w, "")
	CheckErr(err)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	tmpl := path.Join("templates", "new.html")
	t, err := template.ParseFiles(tmpl)
	CheckErr(err)
	err = t.Execute(w, "")
	CheckErr(err)
}

func SavePost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	now := time.Now()
	p := &models.Post{
		Title:   title,
		Created: now,
		Body:    []byte(body),
	}
	core.Insert(p)
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/update/"):]

	// save global oldtitle for late use
	oldtitle = title

	p := core.Query(title)
	renderTemplate(w, p, "edit")
}

func SaveUpdate(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	now := time.Now()
	p := &models.Post{
		Title:   title,
		Created: now,
		Body:    []byte(body),
	}
	// update post by change its title and content
	core.Update(p, oldtitle)
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	var p models.Blogs
	p.Posts = core.QueryAllPost()
	tmpl := path.Join("templates", "lists.html")
	t, err := template.ParseFiles(tmpl)
	CheckErr(err)
	err = t.Execute(w, p)
	CheckErr(err)
}

func ManagePosts(w http.ResponseWriter, r *http.Request) {
	var p models.Blogs
	p.Posts = core.QueryAllPost()
	tmpl := path.Join("templates", "exists.html")
	t, err := template.ParseFiles(tmpl)
	CheckErr(err)
	err = t.Execute(w, p)
	CheckErr(err)
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/"):]
	p := core.Query(title)
	renderTemplate(w, p, "view")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/delete/"):]
	core.Delete(title)
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
}

// cleanup by delete database file and initialize it again
// all exist data will be lost
func CleanUp(w http.ResponseWriter, r *http.Request) {
	core.InitDB()
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
}
