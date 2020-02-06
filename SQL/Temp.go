package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//query, err := db.Query("INSERT INTO test VALUES ( 'TEST', 2 )")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer query.Close()

	selectQuery, err := db.Query("SELECT name, id from test where id = ?", 3)

	for selectQuery.Next() {
		var name *string
		var id *int

		err := selectQuery.Scan(&name, &id)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("Name: %s. Id: %d\n", *name, *id)
	}


	fmt.Println("STarted DB server")

}