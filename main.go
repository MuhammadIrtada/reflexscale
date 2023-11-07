package main

import (
	"os"
	"reflexscale/database/mysql"
	"reflexscale/src/handler"
	"reflexscale/src/repository"
	"reflexscale/src/usecase"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// init DB
	db, err := mysql.Init()
	if err != nil {
		panic(err)
	}

	// run migration
	if os.Getenv("DB_USERNAME") == "root" {
		migration := mysql.Migration{Db: db.DB}
		migration.RunMigration()
	}

	// init repository
	repo := repository.Init(db)

	// init usecase
	uc := usecase.Init(repo)

	// init handler
	rest := handler.Init(uc)
	rest.Run()
}
