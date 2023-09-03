package main

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

func main() {
	sql, params, err := squirrel.Select("u.*").
		From("users as u").
		LeftJoin("profile as p on u.user_id = p.user_id").
		Where(squirrel.Or{
			squirrel.Eq{
				"p.id":   []int{1, 2, 3},
				"p.name": "hogehoge",
			},
		}).
		Where(squirrel.Eq{
			"u.id":   1,
			"u.name": []string{"hoge", "fuga"},
		}).
		ToSql()
	if err != nil {
		panic(err)
	}

	fmt.Println(sql, params)
}
