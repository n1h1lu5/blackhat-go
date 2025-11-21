package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/store")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var (
		ccnum, date, cvv, exp string
		amount                float32
	)
	rows, err := db.Query("select ccnum. date, amount, cvv. exp from transactions")
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ccnum, &date, &amount, &cvv, &exp)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(ccnum, date, amount, cvv, exp)
	}
	if rows.Err() != nil {
		log.Panicln(err)
	}
}
