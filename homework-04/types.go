package main

import "database/sql"

type BlogEntry struct {
	ID string
	Autor string
	Date string
	Header string
	Content string
}

type Blog struct {
	DB *sql.DB
}
