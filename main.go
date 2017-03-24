package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const dbCreds = "host=localhost sslmode=disable port=5432 password=postgres dbname=kperez user=postgres"
const driver = "postgres"

func main() {
	fmt.Println("open db connection")
	db, dbErr := sql.Open(driver, dbCreds)
	if dbErr != nil {
		fmt.Printf("There was an error accessing the db: %v\n", dbErr)
	}
	defer func() {
		// fmt.Println("dropping people table")
		// _, execErr := db.Exec("drop table people")
		// if execErr != nil {
		// 	fmt.Printf("There was an error creating the table: %v\n", execErr)
		// }
		fmt.Println("closing db connection")
		db.Close()
	}()

	// fmt.Println("creating people table")
	// query := `create table people(
	//   id serial primary key,
	//   first_name text,
	//   last_name text,
	//   age int
	//   )`
	// _, execErr := db.Exec(query)
	// if execErr != nil {
	// 	fmt.Printf("There was an error creating the table: %v\n", execErr)
	// 	return
	// }

	fmt.Println("beginning transaction")
	tx, txErr := db.Begin()
	if txErr != nil {
		fmt.Printf("There was an error accessing the db: %v\n", txErr)
	}

	fmt.Println("inserting into people")
	query := "insert into people(first_name, last_name, age) values('kelmer', 'perez', '40')"
	_, execErr := tx.Exec(query)
	if execErr != nil {
		fmt.Printf("There was an error inserting to the db: %v\n", execErr)
		return
	}

	dbErr = tx.Commit()
	if dbErr != nil {
		fmt.Printf("There was an error committing the transaction to the db: %v\n", dbErr)
		return
	}
}
