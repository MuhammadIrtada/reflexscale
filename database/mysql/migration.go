package mysql

import (
	"reflexscale/src/entity"

	"gorm.io/gorm"
)

type Migration struct {
	Db *gorm.DB
}

func (m *Migration) RunMigration() {
	m.Db.AutoMigrate(
		&entity.User{},
		&entity.HasilTest{},
	)
}