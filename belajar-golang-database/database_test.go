package belajar_golang_database

import (
	"database/sql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gunakan DB
}
