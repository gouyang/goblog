package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

var err = errors.New("Open database fail")

func init() {
	db, err = sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.Fatalln(err)
	}
	exist := `select * from blog`
	_, err = db.Exec(exist)
	if err != nil {
		sqlStmt := `CREATE TABLE blog (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, created TIMESTAMP, body BLOB);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (p *post) insert() error {
	now := time.Now().Unix()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmtq, err := db.Prepare("SELECT id FROM blog WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmtq.Close()

	var nid int64
	err = stmtq.QueryRow(now).Scan(&nid)
	if err == nil {
		now = now + 1
	}

	stmt, err := tx.Prepare("INSERT INTO blog VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(now, p.Title, p.Created, p.Body)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (p *post) delete() error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM blog WHERE title = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Title)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (p *post) update(title string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE blog SET title = ?, created = ?, body = ? WHERE title = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Title, p.Created, p.Body, title)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (p *post) query() (np *post, err error) {
	stmt, err := db.Prepare("SELECT title, created, body FROM blog WHERE title = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	np = new(post)
	err = stmt.QueryRow(p.Title).Scan(&np.Title, &np.Created, &np.Body)
	if err != nil {
		return nil, err
	}
	np.Created = np.Created.Local()

	return np, nil
}

func getAllTitles() (titles map[string][]byte, err error) {
	rows, err := db.Query("SELECT title, body FROM blog")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	titles = make(map[string][]byte)
	for rows.Next() {
		var title string
		var body []byte
		rows.Scan(&title, &body)
		titles[title] = body
	}

	return titles, nil
}

func getAllPosts() (p []post) {
	rows, err := db.Query("SELECT title, created, body FROM blog ORDER BY id DESC")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var (
			ttitle   string
			tcreated time.Time
			tbody    []byte
		)
		rows.Scan(&ttitle, &tcreated, &tbody)
		npost := post{Title: ttitle, Created: tcreated.Local(), Body: tbody}
		p = append(p, npost)
	}

	return
}

func cleanup() error {
	defer db.Close()
	stmt := `delete from blog`
	_, err = db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
