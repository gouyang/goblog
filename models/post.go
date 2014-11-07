package models

import (
	"time"
)

type Post struct {
	Title   string
	Body    []byte
	Created time.Time
}
