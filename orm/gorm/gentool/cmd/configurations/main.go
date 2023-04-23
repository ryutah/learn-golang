package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dao",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	dsn := "host=127.0.0.1 user=postgres password=psql dbname=world port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	g.UseDB(db)

	// g.ApplyBasic(
	// 	model.Country{},
	// 	model.Countrylanguage{},
	// 	model.City{},
	// )
	g.GenerateModelAs("city", "MyCity")
	g.ApplyBasic(
		g.GenerateModelAs("city", "MyCity"),
		g.GenerateModel("country"),
	)

	g.Execute()
}
