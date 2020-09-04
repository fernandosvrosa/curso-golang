package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/curso-golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Tiago") // Erro

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

}
