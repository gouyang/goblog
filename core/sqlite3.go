package core

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
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

	sqlStmt := `create table blog (title text not null primary key, body blob);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func SqliteInsert(title string, body []byte) {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	LogFatal(err)

	tx, err := db.Begin()
	LogFatal(err)

	stmt, err := tx.Prepare("insert into blog values(?, ?)")
	LogFatal(err)
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	LogFatal(err)
	tx.Commit()
}

func SqliteQuery() (titles map[string][]byte) {
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
		fmt.Println(title, string(body))
	}
	fmt.Println(titles)
	return titles
}
