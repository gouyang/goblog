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
	r.HandleFunc("/", blog.HomePage)
	r.HandleFunc("/admin", authenticator.Wrap(auth.LoginAdmin))
	r.HandleFunc("/blogs", blog.ListPost)
	r.HandleFunc("/blog/{title}", blog.ViewPost)
	r.HandleFunc("/blog/new/", blog.NewPost)
	r.HandleFunc("/blog/save/{title}", blog.SavePost)
	r.HandleFunc("/blog/save/", blog.SavePost)
	r.HandleFunc("/blog/update/{title}", blog.UpdatePost)
	r.HandleFunc("/blog/saveupdate/", blog.SaveUpdate)
	r.HandleFunc("/blog/saveupdate/{title}", blog.SaveUpdate)
	r.HandleFunc("/blog/delete/{title}", blog.DeletePost)
	r.HandleFunc("/blogs/delete", blog.PostsForDelete)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
