package repository

import (
	"reflexscale/database/mysql"
	"reflexscale/src/entity"
)

type (
	HasilTestInterface interface {
		CreateHasilTest(hasilTest *entity.HasilTest) (*entity.HasilTest, error)
	}

	hasilTest struct {
		db mysql.DB
	}
)

func InitHasilTest(db *mysql.DB) HasilTestInterface {
	return &hasilTest{
		db: *db,
	}
}

func (h *hasilTest) CreateHasilTest(hasilTest *entity.HasilTest) (*entity.HasilTest, error) {
	err := h.db.Create(&hasilTest).Error
	if err != nil {
		return nil, err
	}

	return hasilTest, err
}
