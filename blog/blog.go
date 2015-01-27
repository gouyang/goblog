package blog

import (
	"net/http"
	"time"

	//httpauth "github.com/abbot/go-http-auth"
	. "github.com/ouyanggh/goblog/core"
	db "github.com/ouyanggh/goblog/core/sqlite"
	"github.com/ouyanggh/goblog/models"
)

var oldtitle string

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := "layout"
	p := &models.Post{}
	RenderTemplate(w, tmpl, p)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	tmpl := "new"
	p := &models.Post{}
	RenderTemplate(w, tmpl, p)
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
	db.Insert(p)
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/update/"):]

	// save global oldtitle for late use
	oldtitle = title

	p := db.Query(title)
	tmpl := "edit"
	RenderTemplate(w, tmpl, p)
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
	db.Update(p, oldtitle)
	http.Redirect(w, r, "/blog/"+title, http.StatusFound)
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	var p models.Blogs
	p.Posts = db.QueryAllPost()

	tmpl := "lists"
	RenderTemplate(w, tmpl, p)
}

func ManagePosts(w http.ResponseWriter, r *http.Request) {
	var p models.Blogs
	p.Posts = db.QueryAllPost()

	tmpl := "exists"
	RenderTemplate(w, tmpl, p)
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/"):]
	p := db.Query(title)

	tmpl := "view"
	RenderTemplate(w, tmpl, p)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/blog/delete/"):]
	db.Delete(title)
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
}

// cleanup by delete database file and initialize it again
// all exist data will be lost
func CleanUp(w http.ResponseWriter, r *http.Request) {
	db.Cleanup()
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
}

func Gallerys(w http.ResponseWriter, r *http.Request) {
	tmpl := "gallerys"
	p := &models.Post{}
	RenderTemplate(w, tmpl, p)
}

func AdminPage(w http.ResponseWriter, r *http.Request) {
	tmpl := "admin"
	p := &models.Post{}
	RenderTemplate(w, tmpl, p)
}
