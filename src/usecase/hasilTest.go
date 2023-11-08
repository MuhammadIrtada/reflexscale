package usecase

import (
	"reflexscale/src/entity"
	"reflexscale/src/repository"
)

type (
	HasilTestInterface interface {
		CreateHasilTest(id uint, hasilTestInput *entity.HasilTestInputParam) (*entity.HasilTest, error)
	}

	hasilTest struct {
		hasilTestRepo repository.HasilTestInterface
	}
)

func InitHasilTest(ht repository.HasilTestInterface) HasilTestInterface {
	return &hasilTest{
		hasilTestRepo: ht,
	}
}

func (h *hasilTest) CreateHasilTest(id uint, hasilTestInput *entity.HasilTestInputParam) (*entity.HasilTest, error) {
	hasilTest := &entity.HasilTest{
		UserID:    id,
		Poin:      hasilTestInput.Poin,
	}

	hasilTest, err := h.hasilTestRepo.CreateHasilTest(hasilTest)
	if err != nil {
		return nil, err
	}

	return hasilTest, nil
}