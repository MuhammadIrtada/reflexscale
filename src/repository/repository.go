package repository

import "reflexscale/database/mysql"

type Repository struct {
	User     UserInterface
}

func Init(db *mysql.DB) *Repository {
	return &Repository{
		User:     InitUser(db),
	}
}