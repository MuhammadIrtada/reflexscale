package usecase

import "reflexscale/src/repository"

type Usecase struct {
	User UserInterface
}

func Init(repo *repository.Repository) *Usecase {
	return &Usecase{
		User: InitUser(repo.User),
	}
}