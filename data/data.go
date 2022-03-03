package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Cred struct {
	id       int
	domain   string
	username string
	password string
}

type CredID struct {
	Id     int
	Domain string
}

type Column string
type Value string

const (
	Password Column = "password"
	Username Column = "username"
	Domain   Column = "domain"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./data/gokeep.db")

	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateCredsTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS credentials (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"domain" TEXT,
		"password" TEXT,
		"username" TEXT
	);`

	statement, err := db.Prepare(createTableSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()

	log.Println("Credentials table created.")
}

func InsertCred(domain, username, password string) {
	insertCredSQL := `INSERT INTO credentials (domain, password, username)
	VALUES (?, ?, ?)`

	statement, err := db.Prepare(insertCredSQL)
	if err != nil {
		log.Fatalln(err)
	}
	defer statement.Close()

	_, err = statement.Exec(domain, username, password)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("New credentials saved successfully")
}

func DeleteCred(credId CredID) {
	deleteCredSql := `DELETE FROM credentials WHERE id = ?`

	statement, err := db.Prepare(deleteCredSql)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(credId.Id)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Credentials for %s has been correctly deleted", credId.Domain)
}

func FindCred(domain string) []CredID {
	findCredSql := `SELECT * FROM credentials WHERE domain LIKE '%'||?||'%'`

	rows, err := db.Query(findCredSql, domain)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var domains []CredID
	for rows.Next() {
		var cred Cred
		// Scan() requires pointers to set var values from db colums, 1 pointer per retrieved column
		err = rows.Scan(&cred.id, &cred.domain, &cred.password, &cred.username)
		if err != nil {
			log.Fatalln(err)
		}
		domains = append(domains, CredID{cred.id, cred.domain})
	}

	return domains
}

func FindAllCreds() []CredID {
	findAllCredSql := `SELECT * FROM credentials`

	rows, err := db.Query(findAllCredSql)

	var creds []CredID
	for rows.Next() {
		var cred Cred 
		err := rows.Scan(&cred.id, &cred.domain, &cred.username, &cred.password)

		if err != nil {
			log.Fatalln(err)
		}

		creds = append(creds, CredID{cred.id, cred.domain})
	}

	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	return creds
}

func FindCredById(id int) Cred {
	findCredSql := `SELECT domain, password, username FROM credentials WHERE id = ?`

	var cred Cred

	err := db.QueryRow(findCredSql, id).Scan(&cred.domain, &cred.password, &cred.username)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return cred
}

func UpdateCred(id int, field Column, value Value) {
	setSql := fmt.Sprintf("%s = %q", field, value)

	updateCredSql := fmt.Sprintf(`UPDATE credentials SET %s WHERE id = ?`, setSql)

	statement, err := db.Prepare(updateCredSql)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(id)

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Credentials Updated.")
}
