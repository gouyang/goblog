package main

import (
	"github.com/ouyanggh/goblog/blog"
	"net/http"
)

func main() {
	http.HandleFunc("/", blog.DefaultView)
	http.HandleFunc("/blogs/", blog.BlogList)
	http.HandleFunc("/blog/", blog.BlogView)
	http.HandleFunc("/blog/new/", blog.BlogNew)
	http.HandleFunc("/blog/save/", blog.BlogSave)
	http.HandleFunc("/blog/edit/", blog.BlogEdit)
	http.ListenAndServe(":8080", nil)
}
