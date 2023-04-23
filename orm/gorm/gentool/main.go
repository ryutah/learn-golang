package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ryutah/learn-golang/orm/gorm/gentool/gen/query"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=postgres password=psql dbname=world port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	q := query.Use(db)
	city, err := q.City.WithContext(context.Background()).First()
	if err != nil {
		panic(err)
	}

	v, _ := json.MarshalIndent(city, "", "  ")
	fmt.Println(string(v))
}
