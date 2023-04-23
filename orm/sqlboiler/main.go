package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/ryutah/learn-golang/orm/sqlboiler/gen/model"
)

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=psql dbname=world sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	city, err := model.Cities().One(context.Background(), db)
	if err != nil {
		panic(err)
	}
	v, _ := json.MarshalIndent(city, "", "  ")
	fmt.Println(string(v))
}
