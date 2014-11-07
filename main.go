package main

import (
	"net/http"
	"os"

	httpauth "github.com/abbot/go-http-auth"
	"github.com/gorilla/mux"
	"github.com/ouyanggh/goblog/auth"
	"github.com/ouyanggh/goblog/blog"
	"github.com/ouyanggh/goblog/core"
)

const SQLITEDBFILE = "sqlite3.db"

func main() {
	// create database file
	if _, err := os.Stat(SQLITEDBFILE); err != nil {
		core.InitSqlite3DB()
	}

	authenticator := httpauth.NewBasicAuthenticator("localhost", auth.Secret)

	fs := http.FileServer(http.Dir("static"))

	r := mux.NewRouter()
	r.HandleFunc("/", blog.HomePage)
	r.HandleFunc("/admin", authenticator.Wrap(auth.LoginAdmin))
	r.HandleFunc("/blogs", blog.ListPosts)
	r.HandleFunc("/blog/{title}", blog.ViewPost)
	r.HandleFunc("/blog/new/", blog.NewPost)
	r.HandleFunc("/blog/save/", blog.SavePost)
	r.HandleFunc("/blog/update/{title}", blog.UpdatePost)
	r.HandleFunc("/blog/saveupdate/", blog.SaveUpdate)
	r.HandleFunc("/blogs/delete/", blog.DeleteLists)
	r.HandleFunc("/blog/delete/{title}", blog.DeletePost)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
