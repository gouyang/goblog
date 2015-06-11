package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"github.com/juju/errgo"
)

type post struct {
	Title   string
	Body    []byte
	Created time.Time
}

type posts struct {
	Posts []post
}

type page struct {
	Tmpl string
	Post interface{}
	W    http.ResponseWriter
}

type postContext struct {
	title string
	db    *sql.DB
}

type blogHandler struct {
	*postContext
	h func(*postContext, http.ResponseWriter, *http.Request) error
}

func (bh blogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := bh.h(bh.postContext, w, r)
	if err != nil {
		log.Fatalln(err.(errgo.Locationer).Location())
	}
}

func main() {
	sqlite3db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.Fatalln(err)
	}
	// init table for post
	exist := `select * from blog`
	_, err = sqlite3db.Exec(exist)
	if err != nil {
		sqlStmt := `CREATE TABLE blog (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, created TIMESTAMP, body BLOB);`
		_, err = sqlite3db.Exec(sqlStmt)
		if err != nil {
			log.Fatalln(err)
		}
	}

	bctx := &postContext{title: "", db: sqlite3db}
	fs := http.FileServer(http.Dir("static"))

	r := mux.NewRouter()
	r.Handle("/", blogHandler{bctx, homePage})
	r.Handle("/admin", blogHandler{bctx, adminPage})
	r.Handle("/cleanup", blogHandler{bctx, cleanUp})
	r.Handle("/blogs", blogHandler{bctx, listPosts})
	r.Handle("/gallerys", blogHandler{bctx, gallerys})
	r.Handle("/blog/{title}", blogHandler{bctx, viewPost})
	r.Handle("/blog/new/", blogHandler{bctx, newPost})
	r.Handle("/blog/save/", blogHandler{bctx, savePost})
	r.Handle("/blog/update/{title}", blogHandler{bctx, updatePost})
	r.Handle("/blog/saveupdate/", blogHandler{bctx, saveUpdate})
	r.Handle("/blogs/manage/", blogHandler{bctx, managePosts})
	r.Handle("/blog/delete/{title}", blogHandler{bctx, deletePost})
	r.Handle("/static/", http.StripPrefix("/static", fs))

	authHandler := httpauth.SimpleBasicAuth("admin", "hello")
	http.Handle("/", authHandler(r))
	http.ListenAndServe(":8001", nil)
}
