package mysql

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ouyanggh/goblog/models"
)

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() {
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)
	defer db.Close()

	sqlStmt := `CREATE TABLE blog (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, created TIMESTAMP, body BLOB);`
	_, err = db.Exec(sqlStmt)
	LogFatal(err)
}

func Insert(p *models.Post) {
	now := time.Now().Unix()
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)

	tx, err := db.Begin()
	LogFatal(err)

	stmtq, err := db.Prepare("SELECT id FROM blog WHERE id = ?")
	LogFatal(err)
	defer stmtq.Close()

	var nid int64
	err = stmtq.QueryRow(now).Scan(&nid)
	if err == nil {
		now = now + 1
	}

	stmt, err := tx.Prepare("INSERT INTO blog VALUES(?, ?, ?, ?)")
	LogFatal(err)
	defer stmt.Close()

	_, err = stmt.Exec(now, p.Title, p.Created, p.Body)
	LogFatal(err)
	tx.Commit()
}

func Delete(title string) {
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)

	tx, err := db.Begin()
	LogFatal(err)

	stmt, err := tx.Prepare("DELETE FROM blog WHERE title = ?")
	LogFatal(err)
	defer stmt.Close()

	_, err = stmt.Exec(title)
	LogFatal(err)
	tx.Commit()
}

func Update(np *models.Post, title string) {
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)

	tx, err := db.Begin()
	LogFatal(err)

	stmt, err := tx.Prepare("UPDATE blog SET title = ?, created = ?, body = ? WHERE title = ?")
	LogFatal(err)
	defer stmt.Close()

	_, err = stmt.Exec(np.Title, np.Created, np.Body, title)
	LogFatal(err)
	tx.Commit()
}

func Query(title string) (p *models.Post) {
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)

	stmt, err := db.Prepare("SELECT title, created, body FROM blog WHERE title = ?")
	LogFatal(err)
	defer stmt.Close()

	p = new(models.Post)
	err = stmt.QueryRow(title).Scan(&p.Title, &p.Created, &p.Body)
	LogFatal(err)

	return
}

func QueryAll() (titles map[string][]byte) {
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)
	rows, err := db.Query("SELECT title, body FROM blog")
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

func QueryAllPost() (p []models.Post) {
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)
	rows, err := db.Query("SELECT title, created, body FROM blog")
	LogFatal(err)
	defer rows.Close()
	for rows.Next() {
		var title string
		var created time.Time
		var body []byte
		rows.Scan(&title, &created, &body)
		post := models.Post{Title: title, Created: created, Body: body}
		p = append(p, post)
	}
	return
}
