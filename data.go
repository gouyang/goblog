package main

import (
	"log"
	"net/http"
	"time"

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
