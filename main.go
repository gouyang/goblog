package main

import (
	"net/http"

	//httpauth "github.com/abbot/go-http-auth"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	auth "github.com/nabeken/negroni-auth"
	"github.com/ouyanggh/goblog/blog"
	//"github.com/ouyanggh/goblog/core"
	db "github.com/ouyanggh/goblog/core/sqlite"
)

func main() {
	// Intialize database file or table.
	db.InitDB()

	//authenticator := httpauth.NewBasicAuthenticator("localhost", core.Secret)

	fs := http.FileServer(http.Dir("static"))

	r := mux.NewRouter()
	r.HandleFunc("/", blog.HomePage)
	r.HandleFunc("/admin", blog.AdminPage)
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
	r.Handle("/static/", http.StripPrefix("/static", fs))
	n := negroni.New()
	n.Use(auth.Basic("admin", "123qweP"))
	n.UseHandler(r)
	//n.Run(":8008")
	//r.Handle("/", n)
	http.ListenAndServe(":8080", n)
}
