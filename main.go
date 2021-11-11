package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlQuery()
}

type department struct {
	id          int
	code        string
	description string
	name        string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/simpleapp")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, code, description, name from department")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []department

	for rows.Next() {
		var each = department{}
		var err = rows.Scan(&each.id, &each.code, &each.description, &each.name)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.code + ": " + each.name + ", " + each.description)
	}
}
