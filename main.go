package main

import (
	"database/sql"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
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
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed with error")
	}
}

func main() {
	sqlite3db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Could not open sqlite3 database")
	}
	log.WithFields(log.Fields{
		"database": "./sqlite3.db",
	}).Info("Run the server with database")
	// init table for post
	exist := `select * from blog`
	_, err = sqlite3db.Exec(exist)
	if err != nil {
		sqlStmt := `CREATE TABLE blog (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, created TIMESTAMP, body BLOB);`
		_, err = sqlite3db.Exec(sqlStmt)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("Could not create table blog in db")
		}
	}

	bctx := &postContext{title: "", db: sqlite3db}
	fs := http.FileServer(http.Dir("static"))

	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/admin", adminPage)
	r.Handle("/cleanup", blogHandler{bctx, cleanUp})
	r.Handle("/blogs", blogHandler{bctx, listPosts})
	r.HandleFunc("/gallerys", gallerys)
	r.Handle("/blog/{title}", blogHandler{bctx, viewPost})
	r.HandleFunc("/blog/new/", newPost)
	r.Handle("/blog/save/", blogHandler{bctx, savePost})
	r.Handle("/blog/update/{title}", blogHandler{bctx, updatePost})
	r.Handle("/blog/saveupdate/", blogHandler{bctx, saveUpdate})
	r.Handle("/blogs/manage/", blogHandler{bctx, managePosts})
	r.Handle("/blog/delete/{title}", blogHandler{bctx, deletePost})
	//r.Handle("/static/", http.StripPrefix("/static", fs))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))

	authHandler := httpauth.SimpleBasicAuth("admin", "hello")
	http.Handle("/", authHandler(r))
	log.WithFields(log.Fields{
		"User":   "admin",
		"passwd": "hello",
	}).Info("admin info")
	log.WithFields(log.Fields{
		"url": ":8001",
	}).Info("Server run at")
	http.ListenAndServe(":8001", nil)
}
