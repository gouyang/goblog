package main

import (
	"net/http"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
)

func main() {
	bctx := &postContext{title: ""}
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
