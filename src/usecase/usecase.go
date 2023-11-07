package usecase

import "reflexscale/src/repository"

type Usecase struct {
	User UserInterface
	HasilTest HasilTestInterface
}

func Init(repo *repository.Repository) *Usecase {
	return &Usecase{
		User: InitUser(repo.User),
		HasilTest: InitHasilTest(repo.HasilTest),
	}
}