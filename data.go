package main

import "time"

type post struct {
	Title   string
	Body    []byte
	Created time.Time
}

type blogs struct {
	Posts []post
}
