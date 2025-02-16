package main

import (
	"database/sql"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	defer fmt.Println("Goodbye, World!")
	fmt.Println("I'm going to sleep now...")
	defer fmt.Println("Getting ready to leave...")

	// db connection, file resources, mutexes, etc.
	db, _ := sql.Open("sqlite3", "test.db")
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM test")
	defer rows.Close()

	// open a file and use defer to close it
	f, _ := os.Open("test.txt")
	defer f.Close()

    
}
