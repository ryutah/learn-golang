package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgxutil"
)

type City struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	CountryCode string `db:"countrycode"`
	District    string `db:"district"`
	Population  int    `db:"population"`
}

func main() {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, "postgresql://postgres:psql@127.0.0.1:5432/world?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()

	cities, err := pgxutil.Select[City](ctx, dbpool, "select * from city limit 10", nil, pgx.RowToStructByName[City])
	if err != nil {
		panic(err)
	}

	for _, city := range cities {
		fmt.Println(city)
	}
}
