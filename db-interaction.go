package main

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"termdb/structs"
)

func PrepareSql() {
	db, err := sql.Open("sqlite", "./msg.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable := `create table if not exists tblMessage(
						id INTEGER PRIMARY KEY,
						title VARCHAR NOT NULL,
						message VARCHAR NOT NULL
					)`
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}
}

func InsertMessage(message *string, title *string) {
	db, err := sql.Open("sqlite", "./msg.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := fmt.Sprintf("INSERT INTO tblMessage(title, message) VALUES ('%s', '%s');", *title, *message)
	result, dberr := db.Exec(sql)
	if dberr != nil {
		log.Fatal(dberr)
	}

	rows, errR := result.RowsAffected()
	if errR != nil || rows != int64(1) {
		log.Fatal("Unexpected error while inserting")
	}
}

func GetMessageById(id *int) (*structs.MessageInfo, error) {
	db, err := sql.Open("sqlite", "./msg.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT id, title, message FROM tblMessage WHERE id = '%d';", *id)
	row := db.QueryRow(sql)

	m := &structs.MessageInfo{}
	err = row.Scan(&m.Id, &m.Title, &m.Message)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetAllRecords() (*[]structs.MessageInfo, error) {
	db, err := sql.Open("sqlite", "./msg.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var queryText = fmt.Sprintf("SELECT id, title, message FROM tblMessage;")
	var results []structs.MessageInfo
	row, err := db.Query(queryText)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		m := &structs.MessageInfo{}
		err := row.Scan(&m.Id, &m.Title, &m.Message)

		if err != nil {
			return nil, err
		}
		results = append(results, *m)
	}
	return &results, nil
}
