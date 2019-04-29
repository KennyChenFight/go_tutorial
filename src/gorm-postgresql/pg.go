package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "go"
	sslmode  = "disable"
)

func main() {
	db := connectDB()
	deleteAll(db)
	insertOne(db)
	queryOne(db)
	queryMany(db)
	updateOne(db)
	deleteProduct(db)
}

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func queryMany(db *sql.DB) {
	var id, name, price string

	rows, err := db.Query(" select * from product where name=$1", "TV")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id, name, price)
}

func queryOne(db *sql.DB) {
	var price string
	err := db.QueryRow(" select price from product where name=$1", "TV").Scan(&price)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	fmt.Println(price)
}

func insertOne(db *sql.DB) {
	stmt, err := db.Prepare(" insert into product(id, name, price) values ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(1, "TV", "100")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert one product success")
	}
}

func updateOne(db *sql.DB) {
	stmt, err := db.Prepare("update product set price=$1 where id=$2")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("300", "1")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("update product success")
	}
}

func deleteProduct(db *sql.DB) {
	stmt, err := db.Prepare("delete from product where id=$1")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(1)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("delete product success")
	}
}

func deleteAll(db *sql.DB) {
	_, err := db.Exec("delete from product")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("delete all product success")
	}
}
