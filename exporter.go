package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../KoboReader.sqlite")
	checkErr(err)

	// Get BookIDs from content table
	rows, err := db.Query("SELECT DISTINCT BookID, BookTitle FROM content")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var bookID sql.NullString
		var bookTitle sql.NullString
		err = rows.Scan(&bookID, &bookTitle)
		checkErr(err)
		if bookID.Valid && bookTitle.Valid {
			fmt.Printf("%s: %s\n", bookID.String, bookTitle.String)
		}
	}
	checkErr(rows.Err())

	fmt.Println("")

	bookIDPtr := flag.String("bookID", "11221b5b-04ef-4964-a8db-76396a0c9f77", "Set to the BookID")
	flag.Parse()
	fmt.Printf("Extracting quotes from book: %s\n\n", *bookIDPtr)

	// to get chapter title, we need to go from bookmarks.contentid to content.contentid, then find the item with the next VolumeIndex. The Title of that row is the chapter title.

	query := fmt.Sprintf(`
					SELECT ContentID, ChapterProgress, Text FROM Bookmark
					WHERE VolumeID="%s"
					ORDER BY ContentID, ChapterProgress
					LIMIT 1`, *bookIDPtr)

	rows, err = db.Query(query)
	checkErr(err)
	defer rows.Close()

	var volumeID string
	var contentID string
	var text sql.NullString
	for rows.Next() {
		err = rows.Scan(&volumeID, &contentID, &text)
		checkErr(err)

		if text.Valid {
			trimmedText := strings.Join(strings.Fields(text.String), " ")
			fmt.Println(trimmedText)
			fmt.Println("")
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
