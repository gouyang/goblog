package main

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (pc *postContext) insert(p *post) error {
	now := time.Now().Unix()

	tx, err := pc.db.Begin()
	if err != nil {
		return err
	}
	stmtq, err := pc.db.Prepare("SELECT id FROM blog WHERE id = ?")
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

func (pc *postContext) delete(p *post) error {
	tx, err := pc.db.Begin()
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

func (pc *postContext) update(p *post, title string) error {
	tx, err := pc.db.Begin()
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

func (pc *postContext) query(p *post) (np *post, err error) {
	stmt, err := pc.db.Prepare("SELECT title, created, body FROM blog WHERE title = ?")
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

func (pc *postContext) getAllTitles() (titles map[string][]byte, err error) {
	rows, err := pc.db.Query("SELECT title, body FROM blog")
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

func (pc *postContext) getAllPosts() (p []post) {
	rows, err := pc.db.Query("SELECT title, created, body FROM blog ORDER BY id DESC")
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

func (pc *postContext) cleanup() error {
	defer pc.db.Close()
	stmt := `delete from blog`
	_, err := pc.db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
