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

var db *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "admin:password@/test?parseTime=true")
	LogFatal(err)
	defer db.Close()

	exist := `select * from blog`
	_, err = db.Exec(exist)
	if err != nil {
		sqlStmt := `CREATE TABLE blog (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, created TIMESTAMP, body BLOB);`
		_, err = db.Exec(sqlStmt)
		LogFatal(err)
		utf8 := `ALTER TABLE blog CONVERT TO CHARACTER SET utf8 COLLATE utf8_unicode_ci;`
		_, err = db.Exec(utf8)
		LogFatal(err)
	}
}

func Insert(p *models.Post) {
	now := time.Now().Unix()

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
	stmt, err := db.Prepare("SELECT title, created, body FROM blog WHERE title = ?")
	LogFatal(err)
	defer stmt.Close()

	p = new(models.Post)
	err = stmt.QueryRow(title).Scan(&p.Title, &p.Created, &p.Body)
	if err != nil {
		p.Body = []byte("The post does not exist")
		return
	}
	p.Created = p.Created.Local()

	return
}

func QueryAll() (titles map[string][]byte) {
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
	rows, err := db.Query("SELECT title, created, body FROM blog ORDER BY id DESC")
	LogFatal(err)
	defer rows.Close()
	for rows.Next() {
		var title string
		var created time.Time
		var body []byte
		rows.Scan(&title, &created, &body)
		post := models.Post{Title: title, Created: created.Local(), Body: body}
		p = append(p, post)
	}
	return
}

func Cleanup() {
	defer db.Close()
	exist := `delete from blog`
	_, err := db.Exec(exist)
	log.Fatalln(err)
}
