package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1111 -d mysql

type employee struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:1111@tcp(localhost:3306)/mysqldb?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Printf("[%T][%v]\n\n", db, db)

	// createTable(db)
	// insertRecord(db)
	// queryRecordAll(db)
	// queryRecord(db)
	updateRecord(db)
}

func createTable(db *sql.DB) {
	createCommand := `
	CREATE TABLE employee
	(
		id INT NOT NULL UNIQUE,
		name VARCHAR(20)
	);
	`
	_, err := db.Exec(createCommand)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create table success.\n")
}

func insertRecord(db *sql.DB) {
	insertStmt, err := db.Prepare("INSERT INTO employee (id, name) VALUES (?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	defer insertStmt.Close()
	fmt.Printf("Insert statement: %v\n\n", insertStmt)

	result, err := insertStmt.Exec(102, "mark")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Insert record success: %v\n\n", result)
}

func updateRecord(db *sql.DB) {
	updateStmt, err := db.Prepare("UPDATE employee SET name=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer updateStmt.Close()

	target_id := 100
	e := employee{id: target_id, name: "Robert"}
	updateResult, err := updateStmt.Exec(e.name, e.id)
	if err != nil {
		log.Fatal(err)
	}

	updateRecords, err := updateResult.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%v]\n", updateRecords)
}

func queryRecordAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		e := employee{}
		err := rows.Scan(&e.id, &e.name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[%v][%v]\n", e.id, e.name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func queryRecord(db *sql.DB) {
	rowStmt, err := db.Prepare("SELECT name FROM employee WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	defer rowStmt.Close()

	target_id := 100
	e := employee{id: target_id}
	err = rowStmt.QueryRow(e.id).Scan(&e.name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%v][%v]\n", target_id, e.name)
}
