package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert
	addRow, err := db.Prepare("INSERT INTO user(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	addRow.Exec("横浜　みなと")

	// select
	rows, err := db.Query("SELECT * FROM user") //
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id string
		var name string
		var valid string
		// カラム値の取得
		err := rows.Scan(&id, &name, &valid)
		if err != nil {
			panic(err)
		}
		// 取得結果
		fmt.Println(id, name, valid)
	}

}
