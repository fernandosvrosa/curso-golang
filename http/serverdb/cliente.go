package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Usuario
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega oara funca adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "Desculpa...")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:123456@/curso-golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario
	_ = db.QueryRow("select id, nome from usuarios where id =?", id).Scan(&u.ID, &u.Nome)

	json, _ := json2.Marshal(u)

	w.Header().Set("Content-type", "application/json")
	_, _ = fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456@/curso-golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios ")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		_ = rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json2.Marshal(usuarios)

	w.Header().Set("Content-type", "application/json")
	_, _ = fmt.Fprint(w, string(json))
}

