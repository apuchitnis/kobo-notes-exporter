package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../KoboReader.sqlite")
	checkErr(err)

	rows, err := db.Query("SELECT VolumeID, ContentID, Text FROM Bookmark WHERE VolumeID=\"11221b5b-04ef-4964-a8db-76396a0c9f77\" LIMIT 100")
	checkErr(err)

	var volumeID string
	var contentID string
	var text sql.NullString
	for rows.Next() {
		err = rows.Scan(&volumeID, &contentID, &text)
		checkErr(err)

		// trimmedText := strings.TrimSpace(text.String)
		if text.Valid {
			// fmt.Println(text.String)
			trimmedText := strings.Join(strings.Fields(text.String), " ")
			fmt.Println(trimmedText)
			// fmt.Println(trimmedText)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
