package main

import (
	"net/http"
	"time"
)

func homePage(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	tmpl := "layout"
	p := &post{}
	err := renderTemplate(w, tmpl, p)
	return err
}

func newPost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	tmpl := "new"
	p := &post{}
	err := renderTemplate(w, tmpl, p)
	return err
}

func savePost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	rtitle := r.FormValue("title")
	rbody := r.FormValue("body")
	now := time.Now()
	p := &post{
		Title:   rtitle,
		Created: now,
		Body:    []byte(rbody),
	}
	err := insert(p)
	http.Redirect(w, r, "/blog/"+rtitle, http.StatusFound)
	return err
}

func updatePost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	title := r.URL.Path[len("/blog/update/"):]

	btx.title = title

	p, err := query(title)
	tmpl := "edit"
	err = renderTemplate(w, tmpl, p)
	return err
}

func saveUpdate(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	rtitle := r.FormValue("title")
	rbody := r.FormValue("body")
	now := time.Now()
	p := &post{
		Title:   rtitle,
		Created: now,
		Body:    []byte(rbody),
	}
	err := update(p, btx.title)
	http.Redirect(w, r, "/blog/"+rtitle, http.StatusFound)
	return err
}

func listPosts(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	var p blogs
	p.Posts = queryAllPost()

	tmpl := "lists"
	err = renderTemplate(w, tmpl, p)
	return err
}

func managePosts(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	var p blogs
	p.Posts = queryAllPost()

	tmpl := "exists"
	err = renderTemplate(w, tmpl, p)
	return err
}

func viewPost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	title := r.URL.Path[len("/blog/"):]
	p, err := query(title)

	tmpl := "view"
	err = renderTemplate(w, tmpl, p)
	return err
}

func deletePost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	title := r.URL.Path[len("/blog/delete/"):]
	err := delete(title)
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
	return err
}

// cleanup by delete database file and initialize it again
// all exist data will be lost
func cleanUp(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	err := cleanup()
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
	return err
}

func gallerys(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	tmpl := "gallerys"
	p := &post{}
	err := renderTemplate(w, tmpl, p)
	return err
}

func adminPage(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	tmpl := "admin"
	p := &post{}
	err := renderTemplate(w, tmpl, p)
	return err
}
