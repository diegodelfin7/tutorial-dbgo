package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Equipo struct {
	id     int
	nombre string
	pais   string
}

func main() {
	db, err := sql.Open("mysql", "root:@/db_futbolgo")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM equipo")
	if err != nil {
		panic(err.Error())
	}

	var tmpe Equipo
	var e []Equipo

	for rows.Next() {
		err = rows.Scan(&tmpe.id, &tmpe.nombre, &tmpe.pais)

		if err != nil {
			panic(err.Error())
		}

		e = append(e, tmpe)

	}

	for _, v := range e {
		fmt.Println("valor", v.id)
		fmt.Println("valor", v.nombre)
		fmt.Println("valor", v.pais)
	}

}
