package main

import (
	httpauth "github.com/abbot/go-http-auth"
	"github.com/gorilla/mux"
	"github.com/ouyanggh/goblog/auth"
	"github.com/ouyanggh/goblog/blog"
	"github.com/ouyanggh/goblog/core"
	"net/http"
	"os"
)

const SQLITEDBFILE = "sqlite3.db"

func main() {
	//for test db driver
	if _, err := os.Stat(SQLITEDBFILE); err != nil {
		core.InitSqlite3DB()
	}
	//core.SqliteInsert("test", []byte("test the sqlite3"))
	//core.SqliteQuery()
	authenticator := httpauth.NewBasicAuthenticator("localhost", auth.Secret)

	r := mux.NewRouter()
	r.HandleFunc("/", blog.DefaultView)
	r.HandleFunc("/admin", authenticator.Wrap(auth.LoginAdmin))
	r.HandleFunc("/blogs/", blog.BlogList)
	r.HandleFunc("/blog/{title}", blog.BlogView)
	r.HandleFunc("/blog/new/", blog.BlogNew)
	r.HandleFunc("/blog/save/{title}", blog.BlogSave)
	r.HandleFunc("/blog/save/", blog.BlogSave)
	r.HandleFunc("/blog/edit/{title}", blog.BlogEdit)
	r.HandleFunc("/blog/update/{title}", blog.BlogUpdate)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
