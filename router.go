package main

import (
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t := compileTemplate("layout")
	err := t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed with error")
	}
}

func newPost(w http.ResponseWriter, r *http.Request) {
	t := compileTemplate("new")
	err := t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed with error")
	}
}

func savePost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	rtitle := r.FormValue("title")
	rtitle = strings.TrimSuffix(rtitle, "?")
	rbody := r.FormValue("body")
	now := time.Now()
	p := &post{
		Title:   rtitle,
		Created: now,
		Body:    []byte(rbody),
	}
	err := btx.insert(p)
	log.WithFields(log.Fields{
		"title": rtitle,
	}).Info("New post")
	http.Redirect(w, r, "/blog/"+rtitle, http.StatusFound)
	return err
}

func updatePost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	title := r.URL.Path[len("/blog/update/"):]
	title = strings.TrimSuffix(title, "?")

	btx.title = title
	p := &post{Title: title}
	p, err := btx.query(p)
	pa := &page{Tmpl: "edit", Post: p, W: w}
	err = pa.renderTemplate()
	return err
}

func saveUpdate(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	rtitle := r.FormValue("title")
	rtitle = strings.TrimSuffix(rtitle, "?")
	rbody := r.FormValue("body")
	now := time.Now()
	p := &post{
		Title:   rtitle,
		Created: now,
		Body:    []byte(rbody),
	}
	err := btx.update(p, btx.title)
	log.WithFields(log.Fields{
		"Title":    btx.title,
		"newTitle": rtitle,
	}).Info("Update post")
	http.Redirect(w, r, "/blog/"+rtitle, http.StatusFound)
	return err
}

func viewPost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	title := r.URL.Path[len("/blog/"):]
	title = strings.TrimSuffix(title, "?")
	p := &post{Title: title}
	p, err := btx.query(p)
	pa := &page{Tmpl: "view", Post: p, W: w}
	err = pa.renderTemplate()
	return err
}

func listPosts(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	var p posts
	p.Posts = btx.getAllPosts()

	pa := &page{Tmpl: "lists", Post: p, W: w}
	err := pa.renderTemplate()
	return err
}

func managePosts(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	var p posts
	p.Posts = btx.getAllPosts()

	pa := &page{Tmpl: "exists", Post: p, W: w}
	err := pa.renderTemplate()
	return err
}

func deletePost(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	title := r.URL.Path[len("/blog/delete/"):]
	title = strings.TrimSuffix(title, "?")
	p := &post{}
	p.Title = title
	err := btx.delete(p)
	log.WithFields(log.Fields{
		"title": title,
	}).Info("Delete post")
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
	return err
}

// cleanup by delete database file and initialize it again
// all exist data will be lost
func cleanUp(btx *postContext, w http.ResponseWriter, r *http.Request) error {
	err := btx.cleanup()
	log.Info("Delete database to cleanup all blogs")
	http.Redirect(w, r, "/blogs/manage/", http.StatusFound)
	return err
}

func gallerys(w http.ResponseWriter, r *http.Request) {
	t := compileTemplate("gallerys")
	err := t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed with error")
	}
}

func adminPage(w http.ResponseWriter, r *http.Request) {
	t := compileTemplate("admin")
	err := t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed with error")
	}
}
