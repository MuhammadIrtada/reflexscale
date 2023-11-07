package repository

import "reflexscale/database/mysql"

type Repository struct {
	User     UserInterface
	HasilTest HasilTestInterface
}

func Init(db *mysql.DB) *Repository {
	return &Repository{
		User:     InitUser(db),
		HasilTest: InitHasilTest(db),
	}
}