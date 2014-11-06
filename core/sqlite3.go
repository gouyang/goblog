package core

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ouyanggh/goblog/models"
)

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitSqlite3DB() {
	os.Remove("./sqlite3.db")

	db, err := sql.Open("sqlite3", "./sqlite3.db")
	LogFatal(err)
	defer db.Close()

	sqlStmt := `create table blog (ID integer not null primary key, title text not null, body blob);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

//func SqliteInsert(title string, body []byte) {
func SqliteInsert(p *models.Post) {
	now := time.Now().Unix()
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	LogFatal(err)

	tx, err := db.Begin()
	LogFatal(err)

	stmt, err := tx.Prepare("insert into blog values(?, ?, ?)")
	LogFatal(err)
	defer stmt.Close()

	_, err = stmt.Exec(now, p.Title, p.Body)
	LogFatal(err)
	tx.Commit()
}

func SqliteDelete(title string) {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	LogFatal(err)

	tx, err := db.Begin()
	LogFatal(err)

	stmt, err := tx.Prepare("delete from blog where title = ?")
	LogFatal(err)
	defer stmt.Close()

	_, err = stmt.Exec(title)
	LogFatal(err)
	tx.Commit()
}

func SqliteUpdate(np *models.Post, title string) {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	LogFatal(err)

	tx, err := db.Begin()
	LogFatal(err)

	stmt, err := tx.Prepare("update blog set title = ?, body = ? where title = ?")
	LogFatal(err)
	defer stmt.Close()

	_, err = stmt.Exec(np.Title, np.Body, title)
	LogFatal(err)
	tx.Commit()
}

func SqliteQuery(title string) string {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	LogFatal(err)

	stmt, err := db.Prepare("select body from blog where title = ?")
	LogFatal(err)
	defer stmt.Close()

	var body string
	err = stmt.QueryRow(title).Scan(&body)
	LogFatal(err)

	return body
}

func SqliteQueryAll() (titles map[string][]byte) {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	LogFatal(err)

	rows, err := db.Query("select title, body from blog")
	LogFatal(err)
	defer rows.Close()
	titles = make(map[string][]byte)
	for rows.Next() {
		var title string
		var body []byte
		rows.Scan(&title, &body)
		titles[title] = body
	}
	return titles
}
