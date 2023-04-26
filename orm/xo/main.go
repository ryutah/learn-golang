package main

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ryutah/learn-golang/orm/xo/models"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5432 user=postgres password=psql dbname=world sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var cities []models.City
	if err := db.Select(&cities, "SELECT * FROM city LIMIT 1"); err != nil {
		panic(err)
	}

	v, _ := json.MarshalIndent(cities, "", "  ")
	fmt.Println(string(v))
}
