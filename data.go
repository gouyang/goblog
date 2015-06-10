package main

import (
	"net/http"
	"time"
)

type post struct {
	Title   string
	Body    []byte
	Created time.Time
}

type blogs struct {
	Posts []post
}

type page struct {
	Tmpl string
	Post interface{}
	W    http.ResponseWriter
}
