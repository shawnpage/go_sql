package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(:3307)/test")
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10) // TODO: calibrate
	db.SetMaxIdleConns(10) // TODO: calibrate

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(id bigint(20) unsigned NOT NULL AUTO_INCREMENT, world varchar(50), email varchar(50), latitude float(25), longitude float(25), PRIMARY KEY (id))")
	if err != nil {
		fmt.Println("create table error:")
		log.Fatal(err)
	}

	latitude, longitude := 35.779590, -78.638179
	email := "spage@webassign.net"

	res, err := db.Exec("INSERT INTO test.hello(world, email, latitude, longitude) VALUES(?, ?, ?, ?)", "hello world!", email, latitude, longitude)
	if err != nil {
		//log.Fatal("insert error:")
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserted %d rows", rowCount)

	rows, err := db.Query("SELECT * FROM test.hello")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			s1 string
			s2 sql.NullString
			i1 int
			f1 float64
			f2 float64
		)

		err = rows.Scan(&i1, &s1, &s2, &f1, &f2)

		log.Printf("found row containing %q", s1, s2, i1, f1, f2)
	}

	res, err = db.Exec("DELETE FROM test.hello LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("deleted rows:", rowCnt)

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

}
