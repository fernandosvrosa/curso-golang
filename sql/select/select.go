package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/curso-golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select * from usuarios where id > ?", 0)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		_ = rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
