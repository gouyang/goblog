package main

import (
	"net/http"

	httpauth "github.com/abbot/go-http-auth"
	"github.com/gorilla/mux"
	"github.com/ouyanggh/goblog/blog"
	"github.com/ouyanggh/goblog/core"
	db "github.com/ouyanggh/goblog/core/sqlite"
)

const SQLITEDBFILE = "sqlite3.db"

func main() {
	// create database file
	db.InitDB()

	authenticator := httpauth.NewBasicAuthenticator("localhost", core.Secret)

	fs := http.FileServer(http.Dir("static"))

	r := mux.NewRouter()
	r.HandleFunc("/", blog.HomePage)
	r.HandleFunc("/admin", authenticator.Wrap(blog.LoginAdmin))
	r.HandleFunc("/cleanup", blog.CleanUp)
	r.HandleFunc("/blogs", blog.ListPosts)
	r.HandleFunc("/gallerys", blog.Gallerys)
	r.HandleFunc("/blog/{title}", blog.ViewPost)
	r.HandleFunc("/blog/new/", blog.NewPost)
	r.HandleFunc("/blog/save/", blog.SavePost)
	r.HandleFunc("/blog/update/{title}", blog.UpdatePost)
	r.HandleFunc("/blog/saveupdate/", blog.SaveUpdate)
	r.HandleFunc("/blogs/manage/", blog.ManagePosts)
	r.HandleFunc("/blog/delete/{title}", blog.DeletePost)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
