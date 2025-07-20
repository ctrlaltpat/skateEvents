package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open(("sqlite3"), "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	quickTest(db)
}

func quickTest(db *sql.DB) {
	sts := `
INSERT INTO users (email, name, password) VALUES('ctrlaltpat@gmail.com', 'Patrick', 'password');
INSERT INTO users (email, name, password) VALUES('inevitable@gmail.com', 'Thanos', 'password');
`
	_, err := db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, email FROM users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s %s\n", id, name, email)
	}
}
